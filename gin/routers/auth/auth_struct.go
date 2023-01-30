package auth

import (
	"time"
)

type tokenUserInfo struct {
	RefreshToken string `gorm:"column:refresh_token;type:varchar(255);" json:"refreshed_at"`
}

type TokenBody struct {
	RefreshToken string `json:"refresh_token"`
}

type TokenResponse struct {
	RefreshToken string `json:"refresh_token"` // refresh token
	AccessToken  string `json:"access_token"`  // access token
}

type UserResponse struct {
	Id           string    `json:"-"`
	Type         int       `json:"-"`
	Email        string    `json:"email"` // email 마우스 호버시에만 보이게
	Password     string    `json:"-"`
	Nickname     string    `json:"nickname"`      // nickname
	ProfileImage string    `json:"profile_image"` // google login시 받아오는 profile image
	CreatedAt    time.Time `json:"created_at"`    // 가입일
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}

type RegistBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	NickName string `json:"nickname"`
}

type GoogleBody struct {
	Email        string `json:"email"`
	ProfileImage string `json:"profile_image"`
}
