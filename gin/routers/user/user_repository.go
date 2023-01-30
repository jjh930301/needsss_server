package user

import (
	"encoding/json"

	"github.com/jjh930301/market/routers/auth"
	"github.com/jjh930301/needsss_common/database"
	"github.com/jjh930301/needsss_common/models"
	uuid "github.com/satori/go.uuid"
)

func setInfo(id uuid.UUID, data interface{}) *auth.UserResponse {
	var model auth.UserResponse
	var obj map[string]interface{}
	records, _ := json.Marshal(data)
	json.Unmarshal(records, &obj)
	nickname := obj["nickname"].(string)
	if nickname != "" {
		var exist NicknameBody
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
