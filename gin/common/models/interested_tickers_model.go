package models

import (
	"time"

	"github.com/jjh930301/market/common/constants"
	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type InterestedTickerModel struct {
	UserId       uuid.UUID `gorm:"column:user_id;type:varchar(36);primary_key" json:"-"`
	User         UserModel
	TickerSymbol string `gorm:"column:symbol;type:varchar(12);not null;primary_key" json:"symbol,omitempty"` //index
	Ticker       KrTickerModel
	Date         time.Time       `gorm:"column:date;type:date;not null;primary_key" json:"date,omitempty"`
	DateTime     time.Time       `gorm:"column:date_time;type:datetime;not null" json:"time,omitempty"`
	Type         int             `gorm:"column:type;type:tinyint;not null" json:"-"`
	Name         string          `gorm:"column:name;type:varchar(100);not null;index" json:"name,omitempty"` //index
	Close        decimal.Decimal `gorm:"column:close;type:decimal(11,3);not null" json:"close"`
	Percent      float32         `gorm:"column:percent;type:float;not null" json:"percent"`
	Volume       int64           `gorm:"column:volume;type:bigint;not null" json:"volume"`
	SaleClose    decimal.Decimal `gorm:"column:sale_close;type:decimal(11,3);default:null" json:"sales_close"`
	SaledAt      time.Time       `gorm:"column:saled_at;type:datetime;default:null" json:"saled_at,omitempty"`
}

func (InterestedTickerModel) TableName() string {
	return constants.InterestedTickerTableName
}
