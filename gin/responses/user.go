package responses

import "time"

type TokenResponse struct {
	RefreshToken string `json:"refresh_token"` // refresh token
	AccessToken  string `json:"access_token"`  // access token
}

type UserResponse struct {
	Id           string    `json:"id"`
	Type         int       `json:"-"`
	Email        string    `json:"email"` // email 마우스 호버시에만 보이게
	Password     string    `json:"-"`
	Nickname     string    `json:"nickname"`      // nickname
	ProfileImage string    `json:"profile_image"` // google login시 받아오는 profile image
	CreatedAt    time.Time `json:"created_at"`    // 가입일
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
}
