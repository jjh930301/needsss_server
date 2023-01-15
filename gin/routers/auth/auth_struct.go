package auth

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type tokenUserInfo struct {
	RefreshToken string `gorm:"column:refresh_token;type:varchar(255);" json:"refreshed_at"`
}

type TokenBody struct {
	RefreshToken string `json:"refresh_token"`
}

type TokenResponse struct {
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
}

type UserResponse struct {
	Id           string    `json:"id"`
	Type         int       `json:"-"`
	Email        string    `json:"email"`
	Password     string    `json:"-"`
	Nickname     string    `json:"nickname"`
	ProfileImage string    `json:"profile_image"`
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
	ID           uuid.UUID `gorm:"column:id;type:varchar(36);primary_key" json:"id"`
	Type         int8      `grom:"column:type;type:tinyint;default:0" json:"type"`
	NickName     string    `gorm:"column:nickname;type:varchar(30);default:null" json:"nickname,omitempty"`
	Email        string    `gorm:"column:email;type:varchar(100);not null" json:"email"`
	Password     string    `gorm:"column:password;type:varchar(255);default:null" json:"-"`
	Mobile       string    `gorm:"column:mobile;type:varchar(12);default:null" json:"mobile"`
	ProfileImage string    `gorm:"column:profile_image;type:text;default:null" json:"profile_image"`
	AccessToken  string    `gorm:"-" json:"access_token"`
	RefreshToken string    `gorm:"-" json:"refresh_token"`
}
