package auth

import (
	"fmt"

	"github.com/jjh930301/needsss_common/database"
	"github.com/jjh930301/needsss_common/models"
	"github.com/jjh930301/needsss_common/utils"
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

	return user, nil
}

func googleLogin(body GoogleBody) (*UserResponse, int) {
	user := findByEmail(body.Email)
	if user != nil {
		database.DB.Model(&models.UserModel{}).Where(&models.UserModel{
			Email: body.Email,
		}).Updates(&models.UserModel{
			ProfileImage: body.ProfileImage,
		}).Scan(user)
		return user, 1
	}
	model := &models.UserModel{
		Email:        body.Email,
		ProfileImage: body.ProfileImage,
	}
	database.DB.Create(model)
	newUser := &UserResponse{
		Id:           model.ID.String(),
		Email:        model.Email,
		Nickname:     model.NickName,
		ProfileImage: model.ProfileImage,
		CreatedAt:    model.CreatedAt,
	}
	return newUser, 2
}

func googleRegist(json map[string]interface{}) (*UserResponse, int) {
	user := findByEmail(json["email"].(string))
	if user != nil {
		database.DB.Model(&models.UserModel{}).Where(&models.UserModel{
			Email: json["email"].(string),
		}).Updates(&models.UserModel{
			ProfileImage: json["picture"].(string),
		}).Scan(user)
		return user, 1
	}
	model := &models.UserModel{
		Email:        json["email"].(string),
		ProfileImage: json["picture"].(string),
	}
	database.DB.Create(model)
	newUser := &UserResponse{
		Id:           model.ID.String(),
		Email:        model.Email,
		Nickname:     model.NickName,
		ProfileImage: model.ProfileImage,
		CreatedAt:    model.CreatedAt,
	}
	return newUser, 2
}

func updateRefreshToken(id uuid.UUID, refreshToken string) {
	database.DB.Model(&models.UserModel{}).Where(&models.UserModel{
		ID: id,
	}).Updates(map[string]interface{}{
		"refresh_token": refreshToken,
	})
}

func findRefreshToken(id uuid.UUID) string {
	var user tokenUserInfo
	database.DB.Model(&models.UserModel{}).Select("refresh_token").Where(&models.UserModel{
		ID: id,
	}).Limit(1).Scan(&user)
	return user.RefreshToken
}

func findByEmail(email string) *UserResponse {
	var model *UserResponse
	database.DB.Model(&models.UserModel{}).Where(&models.UserModel{
		Email: email,
	}).Limit(1).Scan(&model)
	fmt.Println(model)
	return model
}
