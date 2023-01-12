package auth

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type UserResponse struct {
	Id           string    `json:"id"`
	Mobile       string    `json:"mobile"`
	CreatedAt    time.Time `json:"created_at"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}

type RegistBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	NickName string `json:"nickname"`
}

type LoginResponse struct {
	ID           uuid.UUID `gorm:"column:id;type:varchar(36);primary_key" json:"-"`
	Type         int8      `grom:"column:type;type:tinyint;default:0" json:"-"`
	NickName     string    `gorm:"column:nickname;type:varchar(30);default:null" json:"nickname,omitempty"`
	Email        string    `gorm:"column:email;type:varchar(100);not null" json:"email"`
	Password     string    `gorm:"column:password;type:varchar(255);not null" json:"-"`
	Mobile       string    `gorm:"column:mobile;type:varchar(12)" json:"mobile"`
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}
