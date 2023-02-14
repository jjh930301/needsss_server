package repositories

import (
	"github.com/jjh930301/needsss/socket/jsons"
	"github.com/jjh930301/needsss_common/database"
	"github.com/jjh930301/needsss_common/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func NewTickerComment(
	body jsons.CommentJson,
) *jsons.TickerComments {
	id, err := uuid.FromString(body.UserId)
	if err != nil {
		return nil
	}
	commentModel := &models.KrTickerCommentModel{
		UserId:       id,
		Comment:      body.Comment,
		TickerSymbol: body.Symbol,
	}
	database.DB.Model(&models.KrTickerCommentModel{}).Create(commentModel)

	var comment jsons.TickerComments
	database.DB.Model(&models.KrTickerCommentModel{}).Preload("User", func(user *gorm.DB) *gorm.DB {
		var commentUser jsons.CommentUser
		return user.Model(&models.UserModel{}).Find(&commentUser)
	}).Where(&models.KrTickerCommentModel{
		ID: commentModel.ID,
	}).Limit(1).First(&comment)
	return &comment
}
