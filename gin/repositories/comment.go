package repositories

import (
	"github.com/jjh930301/needsss/gin/dto"
	"github.com/jjh930301/needsss/gin/responses"
	"github.com/jjh930301/needsss_common/database"
	"github.com/jjh930301/needsss_common/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func GetTickerComments(
	symbol string,
	count int,
	lastTime interface{},
) *[]responses.TickerCommentsResponse {
	var comments []responses.TickerCommentsResponse

	data := database.DB.Model(&models.KrTickerCommentModel{}).Preload("User", func(user *gorm.DB) *gorm.DB {
		var commentUser dto.CommentUser
		return user.Model(&models.UserModel{}).First(&commentUser)
	}).Where(&models.KrTickerCommentModel{
		TickerSymbol: symbol,
	})
	if lastTime == nil {
		data.Limit(count).Find(&comments)
	} else {
		data.Where("created_at > ?", lastTime).Order(
			"created_at desc",
		).Limit(count).Find(&comments)
	}
	return &comments
}

func NewTickerComment(
	id uuid.UUID,
	body dto.NewCommentDto,
) *[]responses.TickerCommentsResponse {
	database.DB.Model(&models.KrTickerCommentModel{}).Create(&models.KrTickerCommentModel{
		UserId:       id,
		Comment:      body.Comment,
		TickerSymbol: body.Symbol,
	})
	var comments []responses.TickerCommentsResponse
	var commentUser dto.CommentUser
	database.DB.Model(&models.KrTickerCommentModel{}).Preload("User", func(user *gorm.DB) *gorm.DB {
		return user.Model(&models.UserModel{}).Find(&commentUser)
	}).Where(&models.KrTickerCommentModel{
		TickerSymbol: body.Symbol,
	}).Order(
		"created_at desc",
	).Limit(30).Find(&comments)
	return &comments
}
