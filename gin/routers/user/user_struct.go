package user

type NicknameBody struct {
	NickName string `gorm:"column:nickname;type:varchar(30);default:null" json:"nickname,omitempty"`
}
