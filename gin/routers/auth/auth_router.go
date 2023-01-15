package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/jjh930301/market/common/constants"
	"github.com/jjh930301/market/common/models"
	"github.com/jjh930301/market/common/res"
	"github.com/jjh930301/market/common/structs"
	"github.com/jjh930301/market/common/utils"
	"github.com/jjh930301/market/routers/logs"
	uuid "github.com/satori/go.uuid"
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
// @Summary 회원가입
// @Description 2001 성공
// @Description 4001 missing bodies
// @Description 4002 Cannot create user
// @Description 4003 Type is not match
// @Accept  json
// @Produce  json
// @Router /auth/regist [post]
// @Param data body RegistBody true "body data".
// @Success 200 {object} UserResponse
func Register(c *gin.Context) {
	var body RegistBody
	if err := c.ShouldBindJSON(&body); err != nil {
		res.BadRequest(c, "Missing Bodies", 4001)
		panic(nil)
	}
	var repoErr error
	var user *models.UserModel
	user, repoErr = regist(&body)
	if repoErr != nil {
		res.BadRequest(c, "Cannot create new user", 4002)
		panic(nil)
	}

	accessToken, _ := utils.CreateToken(user.ID.String(), user.Email, int(user.Type), 0)
	refreshToken, _ := utils.CreateToken(user.ID.String(), user.Email, int(user.Type), 1)

	updateRefreshToken(user.ID, refreshToken)
	response := UserResponse{
		Id:           user.ID.String(),
		Nickname:     user.NickName,
		CreatedAt:    user.CreatedAt,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
	go logs.Log(c.ClientIP(), constants.POST, constants.AuthGroup+constants.Regist)
	res.Created(c, "Successfully", response, 2001)
}

// @Tags auth
// @Summary 로그인
// @Description 2000 성공
// @Accept json
// @Produce json
// @Param email query string false "email"
// @Param key query string false "key"
// @Router /auth/login [get]
// @Success 200 {object} LoginResponse
func Login(c *gin.Context) {
	email := c.Request.URL.Query().Get("email")
	key := c.Request.URL.Query().Get("key")
	if email == "" || key == "" {
		res.BadRequest(c, "Required key and email", 4001)
		panic(nil)
	}
	pw := utils.DecryptBase64(key)
	if pw == "" {
		res.Forbidden(c, "", 4102)
		panic(nil)
	}
	model := findByEmail(email)
	flag := utils.CheckPasswordHash([]byte(pw), []byte(model.Password))
	if !flag {
		res.Forbidden(c, "Password is not match", 4103)
		panic(nil)
	}
	accessToken, _ := utils.CreateToken(model.Id, model.Email, int(model.Type), 0)
	refreshToken, refreshError := utils.CreateToken(model.Id, model.Email, int(model.Type), 1)
	if refreshError != nil {
		res.BadRequest(
			c,
			"Type is not match",
			4003,
		)
		panic(nil)
	}
	uuid, _ := uuid.FromString(model.Id)
	updateRefreshToken(uuid, refreshToken) // wait this function

	model.AccessToken = accessToken
	model.RefreshToken = refreshToken
	go logs.Log(c.ClientIP(), constants.GET, constants.AuthGroup+constants.Login)
	res.Ok(c, "Successfully login", model, 2000)

}
