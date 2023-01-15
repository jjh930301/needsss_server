package models

import (
	"time"

	"github.com/jjh930301/market/common/constants"
)

type LogModel struct {
	ID       int       `gorm:"column:id;type:bigint;" json:"-"`
	Ip       string    `grom:"column:ip;type:varchar(20)" json:"ip"`            // ip
	DateTime time.Time `gorm:"column:date_time;type:datetime" json:"date_time"` // 들어온 날
	Method   int       `gorm:"column:method;type:tinyint" json:"method"`
	EndPoint string    `grom:"column:end_point;type:varchar(50)" json:"end_point"`
}

func (LogModel) TableName() string {
	return constants.LogTabeName
}
