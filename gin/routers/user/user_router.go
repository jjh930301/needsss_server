package user

import (
	"github.com/gin-gonic/gin"
	"github.com/jjh930301/market/common/res"
	"github.com/jjh930301/market/common/structs"
	"github.com/jjh930301/market/common/utils"
	uuid "github.com/satori/go.uuid"
)

// @Tags user
// @Summary 닉네임 변경
// @Description 2000 성공
// @Description 2001 nickname is exists
// @Description 4001 required nickname
// @Description 4101 required login
// @Accept json
// @Produce json
// @Router /user/nickname [put]
// @Param data body NicknameBody true "nickname"
// @Success 200 {object} auth.UserResponse
// @Security BearerAuth
func SetNickname(c *gin.Context) {
	user, userErr := c.Keys["user"].(structs.AuthClaim)
	if !userErr {
		res.Forbidden(c, "Required login", 4101)
		panic(userErr)
	}
	var body NicknameBody

	if err := c.ShouldBindJSON(&body); err != nil {
		res.BadRequest(c, "Required nickname", 4001)
		panic(err)
	}
	uuid, _ := uuid.FromString(user.UserID)
	model := setInfo(uuid, body)
	if model == nil {
		res.Ok(c, "nickname is exists", gin.H{"result": false}, 2001)
		panic(nil)
	}
	accessToken, _ := utils.CreateToken(model.Id, model.Email, model.Type, 0)
	model.AccessToken = accessToken
	res.Ok(c, "Successfully set nickname", model, 2000)

}
