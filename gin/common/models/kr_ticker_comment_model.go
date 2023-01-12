package models

import (
	"time"

	"github.com/jjh930301/market/common/constants"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type KrTickerCommentModel struct {
	ID           uuid.UUID `gorm:"column:id;type:varchar(36);primary_key" json:"id"`
	UserId       uuid.UUID `gorm:"column:user_id;type:varchar(36)" json:"-"`
	User         UserModel
	TickerSymbol string `gorm:"column:symbol;type:varchar(12)" json:"symbol,omitempty"` //index
	Ticker       KrTickerModel
	Comment      string         `grom:"column:comment;type:text;not null" json:"comment"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime;autoCreateTime" json:"created_at" time_format:"unix"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:datetime;autoUpdateTime:milli" json:"updated_at" time_format:"unix"`
	DeletedAt    gorm.DeletedAt `grom:"column:deleted_at;type:datetime;" json:"deleted_at,omitempty"`
}

func (u *KrTickerCommentModel) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4()
	return
}

func (KrTickerCommentModel) TableName() string {
	return constants.KrTickerCommentTableName
}
