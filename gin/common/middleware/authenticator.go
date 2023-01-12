package middleware

import (
	"errors"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/jjh930301/market/common/res"
)

type AuthClaim struct {
	UserID string `json:"id"`    // 유저 ID
	Type   int    `json:"type"`  // 유저 타입
	Email  string `json:"email"` // 유저 메일
	jwt.StandardClaims
}

func extractToken(r *http.Request) string {
	bearToken := r.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, "Bearer ")
	if len(strArr) == 2 {
		return strArr[1]
	}
	return ""
}

func Authenticator(c *gin.Context) {
	claims := AuthClaim{}
	tokenString := extractToken(c.Request)
	key := func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Unexpected Signing Method")
		}
		return []byte(os.Getenv("JWT_ACCESS_SECRET")), nil
	}

	_, err := jwt.ParseWithClaims(tokenString, &claims, key)
	if err != nil {
		res.Unauthorized(c)
		panic(err)
	}
	c.Set("user", claims)
	c.Next()
}
