package dto

type TokenDto struct {
	RefreshToken string `gorm:"column:refresh_token;type:varchar(255);" json:"refresh_token"`
}

type RegistDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Mobile   string `json:"mobile"`
	NickName string `json:"nickname"`
}

type GoogleDto struct {
	Email        string `json:"email"`
	ProfileImage string `json:"profile_image"`
}
