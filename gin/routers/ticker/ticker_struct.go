package ticker

import (
	"time"

	"github.com/shopspring/decimal"
)

type TickerDto struct {
	Ticker string
}

type oneTickerChartModel struct {
	TickerSymbol string          `gorm:"column:ticker_symbol;type:varchar(12)" json:"-"`
	Date         time.Time       `gorm:"column:date;type:date;primary_key;index" json:"date,omitempty"`
	Open         decimal.Decimal `gorm:"column:open;type:decimal(11,3)" json:"open,omitempty"`
	High         decimal.Decimal `gorm:"column:high;type:decimal(11,3)" json:"high,omitempty"`
	Low          decimal.Decimal `gorm:"column:low;type:decimal(11,3)" json:"low,omitempty"`
	Close        decimal.Decimal `gorm:"column:close;type:decimal(11,3)" json:"close,omitempty"`
	Percent      float32         `gorm:"column:percent;type:float;not null" json:"percent,omitempty"`
	Volume       int64           `gorm:"column:volume;type:bigint;not null" json:"volume,omitempty"`
}

type OneTickerResponse struct {
	Symbol    string                `gorm:"column:symbol;type:varchar(12);not null;primary_key" json:"symbol,omitempty"` //index
	Merket    int                   `gorm:"column:market;type:tinyint;not null" json:"market,omitempty"`
	Name      string                `gorm:"column:name;type:varchar(100);not null" json:"name,omitempty"` //index
	MarketCap decimal.Decimal       `gorm:"column:market_cap;type:decimal(17,0);default:null" json:"market_cap,omitempty"`
	Bps       decimal.Decimal       `gorm:"column:bps;type:decimal(11,2);default:0" json:"bps,omitempty"`
	Per       decimal.Decimal       `gorm:"column:per;type:decimal(11,2);default:0" json:"per,omitempty"`
	Pbr       decimal.Decimal       `gorm:"column:pbr;type:decimal(11,2);default:0" json:"pbr,omitempty"`
	Eps       decimal.Decimal       `gorm:"column:eps;type:decimal(11,2);default:0" json:"eps,omitempty"`
	Div       decimal.Decimal       `gorm:"column:div;type:decimal(11,2);default:0" json:"div,omitempty"`
	Dps       decimal.Decimal       `gorm:"column:dps;type:decimal(11,2);default:0" json:"dps,omitempty"`
	Charts    []oneTickerChartModel `gorm:"foreignkey:TickerSymbol" json:"charts,omitempty"`
}
