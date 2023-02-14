package routers

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jjh930301/needsss/gin/dto"
	"github.com/jjh930301/needsss/gin/repositories"
	"github.com/jjh930301/needsss/gin/responses"
	"github.com/jjh930301/needsss_common/constants"
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
// @Param data body dto.TokenDto true "refresh token".
// @Success 200 {object} responses.TokenResponse
func Token(c *gin.Context) {
	var body dto.TokenDto
	if err := c.ShouldBindJSON(&body); err != nil {
		res.BadRequest(c, "Required refresh_token", 4001)
		return
	}
	claims := structs.AuthClaim{}
	_, err := utils.Verification(body.RefreshToken, &claims, 1)
	if err != nil {
		res.Forbidden(c, "Required login", 4101)
		return
	}
	uuid, _ := uuid.FromString(claims.UserID)
	regacyToken := userRepo.FindRefreshTokenById(uuid)
	if regacyToken != body.RefreshToken {
		res.Forbidden(c, "Required login", 4004)
		return
	}
	accessToken, _ := utils.CreateToken(claims.UserID, claims.Email, claims.Type, constants.AccessTokenType)
	refreshToken, _ := utils.CreateToken(claims.UserID, claims.Email, claims.Type, constants.RefreshTokenType)

	userRepo.UpdateRefreshToken(uuid, refreshToken)
	tokens := responses.TokenResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	res.Ok(c, "Successfully token refresh", tokens, 2000)
}

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
// @Success 200 {object} responses.UserResponse
func GoogleIDToken(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		res.BadRequest(c, "Required token", 4004)
		return
	}
	payload, err := idtoken.Validate(context.Background(), token, os.Getenv("GOOGLE_CLIENT_ID"))
	if err != nil {
		res.Forbidden(c, "Cannot verify google id token", 4003)
		return
	}
	user, t := repositories.GoogleRegist(payload.Claims)
	accessToken, _ := utils.CreateToken(user.Id, user.Email, int(user.Type), constants.AccessTokenType)
	refreshToken, refreshError := utils.CreateToken(user.Id, user.Email, int(user.Type), constants.RefreshTokenType)
	if refreshError != nil {
		res.BadRequest(
			c,
			"Type is not match",
			4003,
		)
		return
	}
	uuid, _ := uuid.FromString(user.Id)
	userRepo.UpdateRefreshToken(uuid, refreshToken)

	user.AccessToken = accessToken
	user.RefreshToken = refreshToken

	if t == 1 {
		res.Ok(c, "Successfully login", user, 2000)
		return
	}
	res.Ok(c, "Successfully new user", user, 2001)
}
