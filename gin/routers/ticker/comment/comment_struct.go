package comment

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type TickerCommentsResponse struct {
	UserId    uuid.UUID   `gorm:"column:user_id;type:varchar(36)" json:"id"`
	User      commentUser `gorm:"foreignKey:UserId;references:ID"`
	Comment   string      `grom:"column:comment" json:"comment"` // 코멘트
	UpdatedAt time.Time   `gorm:"column:updated_at" json:"updated_at"`
}

type commentUser struct {
	ID           uuid.UUID `gorm:"column:id;type:varchar(36);primary_key" json:"-"`
	NickName     string    `gorm:"column:nickname;type:varchar(30);default:null" json:"nickname"`    // 닉네임
	ProfileImage string    `gorm:"column:profile_image;type:text;default:null" json:"profile_image"` // 프로필이미지
}

type NewCommentBody struct {
	Code    string `json:"code"`
	Comment string `json:"comment"`
}
