package models

import (
	"time"

	"github.com/jjh930301/market/common/constants"
	"github.com/shopspring/decimal"
)

type KrTickerChartsModel struct {
	Symbol       string `gorm:"column:symbol;type:varchar(12);primary_key;index" json:"symbol,omitempty"` //index
	TickerSymbol string `gorm:"column:ticker_symbol;type:varchar(12);not null" json:"ticker_symbol,omitempty"`
	Ticker       KrTickerModel
	Date         time.Time       `gorm:"column:date;type:date;primary_key;index" json:"date,omitempty"`
	Open         decimal.Decimal `gorm:"column:open;type:decimal(11,3)" json:"open,omitempty"`
	High         decimal.Decimal `gorm:"column:high;type:decimal(11,3)" json:"high,omitempty"`
	Low          decimal.Decimal `gorm:"column:low;type:decimal(11,3)" json:"low,omitempty"`
	Close        decimal.Decimal `gorm:"column:close;type:decimal(11,3)" json:"close,omitempty"`
	Percent      float32         `gorm:"column:percent;type:float;not null" json:"percent,omitempty"`
	Volume       int64           `gorm:"column:volume;type:bigint;not null" json:"volume,omitempty"`
}

func (KrTickerChartsModel) TableName() string {
	return constants.KrTickerChartTableName
}
