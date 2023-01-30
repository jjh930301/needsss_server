package auth

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jjh930301/needsss_common/res"
	"github.com/jjh930301/needsss_common/structs"
	"github.com/jjh930301/needsss_common/utils"
	uuid "github.com/satori/go.uuid"
	idtoken "google.golang.org/api/idtoken"
)

// @Tags auth
// @Summary 토큰
// @Description 2000 성공
// @Description 4001 required refresh_token
// @Description 4101 required login
// @Description 4004 다른 웹에서 로그인되어 있습니다 다시 로그인 해주세요
// @Accept  json
// @Produce  json
// @Router /auth/token [post]
// @Param data body TokenBody true "refresh token".
// @Success 200 {object} TokenResponse
func Token(c *gin.Context) {
	var body TokenBody
	if err := c.ShouldBindJSON(&body); err != nil {
		res.BadRequest(c, "Required refresh_token", 4001)
		panic(err)
	}
	claims := structs.AuthClaim{}
	_, err := utils.Verification(body.RefreshToken, &claims, 1)
	if err != nil {
		res.Forbidden(c, "Required login", 4101)
		panic(err)
	}
	uuid, _ := uuid.FromString(claims.UserID)
	regacyToken := findRefreshToken(uuid)
	if regacyToken != body.RefreshToken {
		res.Forbidden(c, "Required login", 4004)
		panic(nil)
	}
	accessToken, _ := utils.CreateToken(claims.UserID, claims.Email, claims.Type, 0)
	refreshToken, _ := utils.CreateToken(claims.UserID, claims.Email, claims.Type, 1)

	updateRefreshToken(uuid, refreshToken)
	tokens := TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	res.Ok(c, "Successfully token refresh", tokens, 2000)
}

// @Tags auth
// @Summary 구글 로그인
// @Description 2000 성공
// @Accept json
// @Produce json
// @Router /auth/google/login [get]
// func GoogleLogin(c *gin.Context) {
// 	token := utils.GetGoogleToken()
// 	url := utils.GetGoogleLoginURL(token)
// 	c.Redirect(http.StatusMovedPermanently, url)
// }

// @Tags auth
// @Summary 구글 idtoken
// @Description 2000 성공
// @Description 2001 신규 유저 닉네임 입력 화면으로 이동
// @Description 4004 required token
// @Description 4003 not verify id token
// @Accept json
// @Produce json
// @Param token query string false "idtoken"
// @Router /auth/google/token [get]
// @Success 200 {object} UserResponse
func GoogleIDToken(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		res.BadRequest(c, "Required token", 4004)
		panic(nil)
	}
	payload, err := idtoken.Validate(context.Background(), token, os.Getenv("GOOGLE_CLIENT_ID"))
	if err != nil {
		res.Forbidden(c, "Cannot verify google id token", 4003)
		panic(err)
	}
	user, t := googleRegist(payload.Claims)
	accessToken, _ := utils.CreateToken(user.Id, user.Email, int(user.Type), 0)
	refreshToken, refreshError := utils.CreateToken(user.Id, user.Email, int(user.Type), 1)
	if refreshError != nil {
		res.BadRequest(
			c,
			"Type is not match",
			4003,
		)
		panic(nil)
	}
	uuid, _ := uuid.FromString(user.Id)
	updateRefreshToken(uuid, refreshToken)

	user.AccessToken = accessToken
	user.RefreshToken = refreshToken

	if t == 1 {
		res.Ok(c, "Successfully login", user, 2000)
		panic(nil)
	}
	res.Ok(c, "Successfully new user", user, 2001)
}
