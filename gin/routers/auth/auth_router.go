package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/jjh930301/market/common/models"
	"github.com/jjh930301/market/common/res"
	"github.com/jjh930301/market/common/utils"
)

// @Tags auth
// @Summary 회원가입
// @Description 2001 성공 \n 4001 missing bodies \n 4002 Cannot create user \n 4003 Type is not match
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
	refreshToken, refreshError := utils.CreateToken(user.ID.String(), user.Email, int(user.Type), 1)
	if refreshError != nil {
		res.BadRequest(
			c,
			"Type is not match",
			4003,
		)
		panic(nil)
	}
	updateRefreshToken(user.ID, refreshToken)
	response := UserResponse{
		Id:           user.ID.String(),
		Mobile:       user.Mobile,
		CreatedAt:    user.CreatedAt,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}
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
	var model LoginResponse = findOne(email, pw)
	flag := utils.CheckPasswordHash([]byte(pw), []byte(model.Password))
	if !flag {
		res.Forbidden(c, "Password is not match", 4103)
		panic(nil)
	}
	accessToken, _ := utils.CreateToken(model.ID.String(), model.Email, int(model.Type), 0)
	refreshToken, refreshError := utils.CreateToken(model.ID.String(), model.Email, int(model.Type), 1)
	if refreshError != nil {
		res.BadRequest(
			c,
			"Type is not match",
			4003,
		)
		panic(nil)
	}
	updateRefreshToken(model.ID, refreshToken) // wait this function

	model.AccessToken = accessToken
	model.RefreshToken = refreshToken

	res.Ok(c, "Successfully login", model, 2000)

}
