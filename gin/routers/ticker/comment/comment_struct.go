package comment

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type TickerCommentsResponse struct {
	UserId    uuid.UUID   `gorm:"column:user_id;type:varchar(36)" json:"-"`
	User      commentUser `gorm:"foreignKey:UserId;references:ID"`
	Comment   string      `grom:"column:comment" json:"comment"`
	UpdatedAt time.Time   `gorm:"column:updated_at" json:"updated_at"`
}

type commentUser struct {
	ID           uuid.UUID `gorm:"column:id;type:varchar(36);primary_key" json:"-"`
	NickName     string    `gorm:"column:nickname;type:varchar(30);default:null" json:"nickname"`
	ProfileImage string    `gorm:"column:profile_image;type:text;default:null" json:"profile_image"`
}

type NewCommentBody struct {
	Code    string `json:"code"`
	Comment string `json:"comment"`
}
