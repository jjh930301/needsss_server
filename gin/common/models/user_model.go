package models

import (
	"time"

	"github.com/jjh930301/market/common/constants"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type UserModel struct {
	ID           uuid.UUID               `gorm:"column:id;type:varchar(36);primary_key" json:"id"`
	Type         int8                    `grom:"column:type;type:tinyint;default:0" json:"type"`
	NickName     string                  `gorm:"column:nickname;type:varchar(30);default:null" json:"nickname,omitempty"`
	Email        string                  `gorm:"column:email;type:varchar(100);not null;unique" json:"email"`
	Password     string                  `gorm:"column:password;type:varchar(255);default:null" json:"-"`
	Mobile       string                  `gorm:"column:mobile;type:varchar(12);default:null" json:"mobile"`
	ProfileImage string                  `gorm:"column:profile_image;type:text;default:null" json:"profile_image"`
	RefreshToken string                  `gorm:"column:refresh_token;type:varchar(255);" json:"refreshed_at"`
	CreatedAt    time.Time               `gorm:"column:created_at;type:datetime;autoCreateTime" json:"created_at" time_format:"unix"`
	UpdatedAt    time.Time               `gorm:"column:updated_at;type:datetime;autoUpdateTime:milli" json:"updated_at" time_format:"unix"`
	DeletedAt    gorm.DeletedAt          `grom:"column:deleted_at;type:datetime;" json:"deleted_at,omitempty"`
	Interestes   []InterestedTickerModel `gorm:"foreignKey:UserId" json:"user_list,omitempty"`
}

func (u *UserModel) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = uuid.NewV4()
	return
}

func (UserModel) TableName() string {
	return constants.UserTableName
}
