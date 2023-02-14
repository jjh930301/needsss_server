package repositories

import (
	"encoding/json"
	"strings"

	"github.com/jjh930301/needsss/gin/dto"
	"github.com/jjh930301/needsss/gin/responses"
	"github.com/jjh930301/needsss_common/database"
	"github.com/jjh930301/needsss_common/models"
	"github.com/jjh930301/needsss_common/utils"
	uuid "github.com/satori/go.uuid"
)

type UserRepository struct{}

// Deprecated
func Regist(body *dto.RegistDto) (*models.UserModel, error) {
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

// Deprecated
func GoogleLogin(body dto.GoogleDto) (*responses.UserResponse, int) {
	user := FindByEmail(body.Email)
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
	newUser := &responses.UserResponse{
		Id:           model.ID.String(),
		Email:        model.Email,
		Nickname:     model.NickName,
		ProfileImage: model.ProfileImage,
		CreatedAt:    model.CreatedAt,
	}
	return newUser, 2
}

func GoogleRegist(json map[string]interface{}) (*responses.UserResponse, int) {
	user := FindByEmail(json["email"].(string))
	if user != nil {
		database.DB.Model(&models.UserModel{}).Where(&models.UserModel{
			Email: json["email"].(string),
		}).Updates(&models.UserModel{
			ProfileImage: json["picture"].(string),
		}).Scan(user)
		return user, 1
	}
	nickname := strings.Split(json["email"].(string), "@")[0]
	model := &models.UserModel{
		Email:        json["email"].(string),
		ProfileImage: json["picture"].(string),
		NickName:     nickname,
	}
	database.DB.Create(model)
	newUser := &responses.UserResponse{
		Id:           model.ID.String(),
		Email:        model.Email,
		Nickname:     model.NickName,
		ProfileImage: model.ProfileImage,
		CreatedAt:    model.CreatedAt,
	}
	return newUser, 2
}

func (u UserRepository) UpdateRefreshToken(id uuid.UUID, refreshToken string) {
	database.DB.Model(&models.UserModel{}).Where(&models.UserModel{
		ID: id,
	}).Updates(map[string]interface{}{
		"refresh_token": refreshToken,
	})
}

func (u UserRepository) FindRefreshTokenById(id uuid.UUID) string {
	var user dto.TokenDto
	database.DB.Model(&models.UserModel{}).Select("refresh_token").Where(&models.UserModel{
		ID: id,
	}).Limit(1).Scan(&user)
	return user.RefreshToken
}

func FindByEmail(email string) *responses.UserResponse {
	var model *responses.UserResponse
	database.DB.Model(&models.UserModel{}).Where(&models.UserModel{
		Email: email,
	}).Limit(1).Scan(&model)
	return model
}

func (u UserRepository) SetInfo(id uuid.UUID, data interface{}) *responses.UserResponse {
	var model responses.UserResponse
	var obj map[string]interface{}
	records, _ := json.Marshal(data)
	json.Unmarshal(records, &obj)
	nickname := obj["nickname"].(string)
	if nickname != "" {
		var exist dto.NicknameDto
		database.DB.Model(&models.UserModel{}).Where(&models.UserModel{
			NickName: nickname,
		}).Limit(1).Scan(&exist)
		if exist.NickName != "" {
			return nil
		}
		database.DB.Model(&models.UserModel{}).Where(&models.UserModel{
			ID: id,
		}).Updates(&models.UserModel{
			NickName: nickname,
		}).Scan(&model)
		return &model
	}
	return &model
}
