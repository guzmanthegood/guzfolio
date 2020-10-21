package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"guzfolio/model"

	"github.com/dgrijalva/jwt-go"
)

type User struct {
	UserID		uint
	UserEmail 	string
	IsAdmin		bool
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// get authorization token "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoxLCJ..."
		tokenString, err := getAuthorizationTokenString(r)
		if err != nil {
			http.Error(w, formattedErrorMessage(err.Error()), http.StatusUnauthorized)
			return
		}

		// parse jwt token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})
		if err != nil {
			http.Error(w, formattedErrorMessage(err.Error()), http.StatusUnauthorized)
			return
		}

		// parse token claims into auth user
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, formattedErrorMessage("not valid jwt token"), http.StatusUnauthorized)
			return
		}

		user := &User{
			UserID:    uint(claims["user_id"].(float64)),
			UserEmail: claims["user_email"].(string),
			IsAdmin:   claims["is_admin"].(bool),
		}

		ctx := context.WithValue(r.Context(), contextKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getAuthorizationTokenString(r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return "", errors.New("authorization header required")
	}

	authorizationHeaderParts := strings.SplitN(authHeader, " ", 2)
	if len(authorizationHeaderParts) != 2 || authorizationHeaderParts[0] != "Bearer" {
		return "", errors.New("authorization header format must be Bearer")
	}

	return authorizationHeaderParts[1], nil
}

func formattedErrorMessage(msg string) string {
	e, _ := json.Marshal(model.GraphQLError(msg, "auth middleware"))
	return string(e)
}

type contextKeyType struct{ name string }

var contextKey = contextKeyType{"userContext"}

func ContextAuthUser(ctx context.Context) *User {
	return ctx.Value(contextKey).(*User)
}