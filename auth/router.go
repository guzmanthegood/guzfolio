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
	r.Handle("/register", register(ds))
	r.Handle("/login", login(ds))
	return r
}

func register(ds datastore.DataStore) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// create user
		input := model.CreateUserInput{
			Email:    getURLParam(r, "email"),
			Name:     getURLParam(r, "password"),
			Password: getURLParam(r, "name"),
		}

		user, err := ds.CreateUser(input)
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

		render.PlainText(w, r, token.AccessToken)
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

		render.PlainText(w, r, token.AccessToken)
	})
}

func getURLParam(r *http.Request, key string) string {
	value, ok := r.URL.Query()[key]
	if !ok {
		return ""
	}
	return strings.Join(value,",")
}