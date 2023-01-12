package utils

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func CreateToken(
	id string,
	email string,
	t int,
	tokenType int,
) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["id"] = id
	atClaims["type"] = t
	atClaims["email"] = email
	// 0 == access token
	if tokenType == 0 {
		var err error
		atClaims["exp"] = time.Now().Add(time.Minute * 60).Unix()
		at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
		token, err := at.SignedString([]byte(os.Getenv("JWT_ACCESS_SECRET")))
		if err != nil {
			return "", err
		}
		return token, nil
	}
	// 1 == refresh token
	if tokenType == 1 {
		var err error
		atClaims["exp"] = time.Now().Add(time.Minute * 24).Unix()
		at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
		token, err := at.SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET")))
		if err != nil {
			return "", err
		}
		return token, nil
	}
	return "", nil
}
