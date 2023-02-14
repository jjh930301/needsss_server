package dto

import uuid "github.com/satori/go.uuid"

type CommentUser struct {
	ID           uuid.UUID `gorm:"column:id;type:varchar(36);primary_key" json:"-"`
	NickName     string    `gorm:"column:nickname;type:varchar(30);default:null" json:"nickname"`    // 닉네임
	ProfileImage string    `gorm:"column:profile_image;type:text;default:null" json:"profile_image"` // 프로필이미지
}

type NewCommentDto struct {
	Symbol  string `json:"symbol"`
	Comment string `json:"comment"`
}
