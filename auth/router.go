package auth

import (
	"net/http"
	"strings"

	"guzfolio/datastore"
	"guzfolio/model"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"golang.org/x/crypto/bcrypt"
)

func Router(ds datastore.DataStore) http.Handler {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Handle("/register", register(ds))
	r.Handle("/login", login(ds))
	return r
}

func register(ds datastore.DataStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get url params
		email := getURLParam(r, "email")
		password := getURLParam(r, "password")
		name := getURLParam(r, "name")

		// secure password
		bytePassword := []byte(password)
		passwordHash, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
		if err != nil {
			render.PlainText(w, r, err.Error())
			return
		}

		// create user
		input := model.RegisterInput{
			Email:    email,
			Name:     name,
			Password: string(passwordHash),
		}

		user, err := ds.RegisterUser(input)
		if err != nil {
			render.PlainText(w, r, err.Error())
			return
		}

		// generate new access token
		token, err := generateToken(user.ID, user.Email, false)
		if err != nil {
			render.PlainText(w, r, err.Error())
			return
		}

		authResponse := &model.AuthResponse{
			AuthToken: &model.AuthToken{
				AccessToken: token.AccessToken,
				ExpiredAt:   token.ExpiredAt,
			},
			User: user,
		}

		render.JSON(w, r, authResponse)
	})
}

func login(ds datastore.DataStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email := getURLParam(r, "email")
		password := getURLParam(r, "password")

		// get user
		user, err := ds.GetUserByEmail(email)
		if err != nil {
			render.PlainText(w, r, err.Error())
			return
		}

		// compare password
		bytePassword := []byte(password)
		byteHashedPassword := []byte(user.Password)
		if err := bcrypt.CompareHashAndPassword(byteHashedPassword, bytePassword); err != nil {
			render.PlainText(w, r, err.Error())
			return
		}

		// generate new access token
		token, err := generateToken(user.ID, user.Email, false)
		if err != nil {
			render.PlainText(w, r, err.Error())
			return
		}

		authResponse := &model.AuthResponse{
			AuthToken: &model.AuthToken{
				AccessToken: token.AccessToken,
				ExpiredAt:   token.ExpiredAt,
			},
			User: user,
		}

		render.JSON(w, r, authResponse)
	})
}

func getURLParam(r *http.Request, key string) string {
	value, ok := r.URL.Query()[key]
	if !ok {
		return ""
	}
	return strings.Join(value,",")
}