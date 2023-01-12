package auth

import (
	"github.com/jjh930301/market/common/database"
	"github.com/jjh930301/market/common/models"
	"github.com/jjh930301/market/common/utils"
	uuid "github.com/satori/go.uuid"
)

func regist(body *RegistBody) (*models.UserModel, error) {
	var err error
	pw, err := utils.HashPassword(body.Password)
	if err != nil {
		return nil, err
	}
	user := &models.UserModel{
		Email:    body.Email,
		Password: pw,
		Mobile:   body.Mobile,
		NickName: body.NickName,
	}
	newUser := database.DB.Create(user)
	if newUser == nil {
		return nil, err
	}

	return user, err
}

func updateRefreshToken(id uuid.UUID, refreshToken string) {
	database.DB.Model(&models.UserModel{}).Where(&models.UserModel{
		ID: id,
	}).Updates(map[string]interface{}{
		"refresh_token": refreshToken,
	})
}

func findOne(email string, pw string) LoginResponse {
	var model LoginResponse
	database.DB.Model(&models.UserModel{}).Where(&models.UserModel{
		Email: email,
	}).Scan(&model)

	return model
}
