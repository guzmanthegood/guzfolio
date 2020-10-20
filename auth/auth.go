package auth

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Token struct {
	AccessToken string    	`json:"accessToken"`
	ExpiredAt   time.Time 	`json:"expiredAt"`
}

type Claims struct {
	UserID		uint		`json:"user_id"`
	UserEmail 	string 		`json:"user_email"`
	IsAdmin		bool		`json:"is_admin"`
	jwt.StandardClaims
}

func generateToken(userID uint, userEmail string, isAdmin bool) (*Token, error) {
	expiredAt := time.Now().Add(time.Hour * 24 * 1) // a day

	// create the claims
	claims := Claims{
		UserID:         userID,
		UserEmail:      userEmail,
		IsAdmin:        isAdmin,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiredAt.Unix(),
			IssuedAt:  time.Now().Unix(),
			Issuer:    "guzfolio",
		},
	}

	// generate token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// sign token
	accessToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken: accessToken,
		ExpiredAt:   expiredAt,
	}, nil
}