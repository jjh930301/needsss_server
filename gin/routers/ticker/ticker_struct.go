package ticker

import (
	"time"

	"github.com/shopspring/decimal"
)

type TickerDto struct {
	Ticker string
}

type OneTickerChartResponse struct {
	TickerSymbol string          `gorm:"column:ticker_symbol;type:varchar(12)" json:"-"`
	Date         time.Time       `gorm:"column:date;type:date;primary_key;index" json:"date"`
	Open         decimal.Decimal `gorm:"column:open;type:decimal(11,3)" json:"open"`
	High         decimal.Decimal `gorm:"column:high;type:decimal(11,3)" json:"high"`
	Low          decimal.Decimal `gorm:"column:low;type:decimal(11,3)" json:"low"`
	Close        decimal.Decimal `gorm:"column:close;type:decimal(11,3)" json:"close"`
	Percent      string          `gorm:"column:percent;type:float;not null" json:"percent"`
	Volume       string          `gorm:"column:volume;type:bigint;not null" json:"volume"`
}

type OneTickerResponse struct {
	Symbol         string                   `gorm:"column:symbol;type:varchar(12);not null;primary_key" json:"symbol"` //index
	Merket         int                      `gorm:"column:market;type:tinyint;not null" json:"market"`
	Name           string                   `gorm:"column:name;type:varchar(100);not null" json:"name"` //index
	Bps            decimal.Decimal          `gorm:"column:bps;type:decimal(11,2);default:0" json:"bps"`
	Per            decimal.Decimal          `gorm:"column:per;type:decimal(11,2);default:0" json:"per"`
	Pbr            decimal.Decimal          `gorm:"column:pbr;type:decimal(11,2);default:0" json:"pbr"`
	Eps            decimal.Decimal          `gorm:"column:eps;type:decimal(11,2);default:0" json:"eps"`
	Div            decimal.Decimal          `gorm:"column:div;type:decimal(11,2);default:0" json:"div"`
	Dps            decimal.Decimal          `gorm:"column:dps;type:decimal(11,2);default:0" json:"dps"`
	Sector         string                   `gorm:"column:sector;type:varchar(150)" json:"sector"`
	Industry       string                   `gorm:"column:industry;type:varchar(150)" json:"industry"`
	ListingDate    time.Time                `gorm:"column:listing_date;default:null" json:"listing_date"`
	SettleMonth    string                   `gorm:"column:settle_month;type:varchar(10)" json:"settle_month"`
	Representative string                   `gorm:"column:representative;type:varchar(100)" json:"representative"`
	Homepage       string                   `gorm:"column:homepage;type:varchar(200)" json:"homepage"`
	MarketCap      decimal.Decimal          `gorm:"column:market_cap;type:decimal(17,0);default:null" json:"market_cap"`
	Chart          []OneTickerChartResponse `gorm:"foreignkey:TickerSymbol" json:"chart"`
}

type SearchTickerResponse struct {
	Symbol   string `gorm:"column:symbol;type:varchar(12);not null;primary_key" json:"symbol"` //index
	Merket   int    `gorm:"column:market;type:tinyint;not null" json:"market"`
	Name     string `gorm:"column:name;type:varchar(100);not null" json:"name"` //index
	Homepage string `gorm:"column:homepage;type:varchar(200)" json:"homepage"`
}
