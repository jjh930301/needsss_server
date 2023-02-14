package dto

type NicknameDto struct {
	NickName string `gorm:"column:nickname;type:varchar(30);default:null" json:"nickname,omitempty"`
}
