package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jjh930301/needsss/gin/dto"
	"github.com/jjh930301/needsss_common/constants"
	"github.com/jjh930301/needsss_common/res"
	"github.com/jjh930301/needsss_common/structs"
	"github.com/jjh930301/needsss_common/utils"
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
// @Param data body dto.NicknameDto true "nickname"
// @Success 200 {object} responses.UserResponse
// @Security BearerAuth
func SetNickname(c *gin.Context) {
	user, userErr := c.Keys["user"].(structs.AuthClaim)
	if !userErr {
		res.Forbidden(c, "Required login", 4101)
		return
	}
	var body dto.NicknameDto

	if err := c.ShouldBindJSON(&body); err != nil {
		res.BadRequest(c, "Required nickname", 4001)
		return
	}
	uuid, _ := uuid.FromString(user.UserID)
	model := userRepo.SetInfo(uuid, body)
	if model == nil {
		res.Ok(c, "nickname is exists", gin.H{"result": false}, 2001)
		return
	}
	accessToken, _ := utils.CreateToken(model.Id, model.Email, model.Type, constants.AccessTokenType)
	refreshToken, _ := utils.CreateToken(model.Id, model.Email, model.Type, constants.RefreshTokenType)
	model.AccessToken = accessToken
	model.RefreshToken = refreshToken
	res.Ok(c, "Successfully set nickname", model, 2000)

}
