package jsons

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type CommentJson struct {
	Symbol  string `json:"symbol"`
	Comment string `json:"comment"`
	UserId  string `json:"user_id"`
}

type TickerComments struct {
	UserId    uuid.UUID   `gorm:"column:user_id;type:varchar(36)" json:"-"`
	User      CommentUser `gorm:"foreignKey:UserId;references:ID" json:"user"`
	CommentId string      `gorm:"column:id" json:"comment_id"`         // 코멘트 uuid 수정이나 삭제시 필요할 수도 있습니다.
	Comment   string      `gorm:"column:comment" json:"comment"`       // 코멘트
	Symbol    string      `gorm:"column:symbol" json:"symbol"`         // 메세지를 뿌려줄 symbol
	CreatedAt time.Time   `gorm:"column:created_at" json:"created_at"` // 코멘트 등록일시
}

type CommentUser struct {
	ID           uuid.UUID `gorm:"column:id;type:varchar(36);primary_key" json:"-"`
	NickName     string    `gorm:"column:nickname;type:varchar(30);default:null" json:"nickname"`    // 닉네임
	ProfileImage string    `gorm:"column:profile_image;type:text;default:null" json:"profile_image"` // 프로필이미지
}

type ContextJson struct {
	UserId   string `json:"user_id"`
	FcmToken string `json:"fcm_token"`
}
