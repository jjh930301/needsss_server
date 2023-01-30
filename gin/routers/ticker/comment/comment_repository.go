package comment

import (
	"time"

	"github.com/jjh930301/needsss_common/database"
	"github.com/jjh930301/needsss_common/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func getTickerComments(
	code string,
	lastTime time.Time,
	count int,
) *[]TickerCommentsResponse {
	var comments []TickerCommentsResponse
	var commentUser commentUser
	database.DB.Model(&models.KrTickerCommentModel{}).Preload("User", func(user *gorm.DB) *gorm.DB {
		return user.Model(&models.UserModel{}).Find(&commentUser)
	}).Where(&models.KrTickerCommentModel{
		TickerSymbol: code,
	}).Where("created_at < ?", lastTime).Order(
		"created_at desc",
	).Limit(30).Find(&comments)
	return &comments
}

func newTickerComment(
	id uuid.UUID,
	body NewCommentBody,
) *[]TickerCommentsResponse {
	database.DB.Model(&models.KrTickerCommentModel{}).Create(&models.KrTickerCommentModel{
		UserId:       id,
		Comment:      body.Comment,
		TickerSymbol: body.Code,
	})
	var comments []TickerCommentsResponse
	var commentUser commentUser
	database.DB.Model(&models.KrTickerCommentModel{}).Preload("User", func(user *gorm.DB) *gorm.DB {
		return user.Model(&models.UserModel{}).Find(&commentUser)
	}).Where(&models.KrTickerCommentModel{
		TickerSymbol: body.Code,
	}).Order(
		"created_at desc",
	).Limit(30).Find(&comments)
	return &comments
}
