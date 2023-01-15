package utils

import (
	"errors"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jjh930301/market/common/structs"
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
		atClaims["exp"] = time.Now().Add(time.Minute * 60 * 24 * 7).Unix()
		at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
		token, err := at.SignedString([]byte(os.Getenv("JWT_REFRESH_SECRET")))
		if err != nil {
			return "", err
		}
		return token, nil
	}
	return "", nil
}

// t = 0 access
// t = 1 refresh
func Verification(
	tokenString string,
	claims *structs.AuthClaim,
	t int,
) (*jwt.Token, error) {
	key := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		if t == 0 {
			return []byte(os.Getenv("JWT_ACCESS_SECRET")), nil
		} else {
			return []byte(os.Getenv("JWT_REFRESH_SECRET")), nil
		}
	}

	return jwt.ParseWithClaims(tokenString, claims, key)
}
