package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jjh930301/market/common/res"
	"github.com/jjh930301/market/common/structs"
	"github.com/jjh930301/market/common/utils"
)

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
	claims := structs.AuthClaim{}
	tokenString := extractToken(c.Request)

	_, err := utils.Verification(tokenString, &claims, 0)
	if err != nil {
		res.Unauthorized(c)
		panic(err)
	}
	c.Set("user", claims)
	c.Next()
}
