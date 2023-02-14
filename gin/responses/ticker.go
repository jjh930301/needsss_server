package responses

import (
	"time"

	"github.com/shopspring/decimal"
)

type OneTickerChartResponse struct {
	TickerSymbol string          `gorm:"column:ticker_symbol;type:varchar(12)" json:"-"`
	Date         time.Time       `gorm:"column:date;type:date;primary_key;index" json:"date"` // 날짜
	Open         decimal.Decimal `gorm:"column:open;type:decimal(11,3)" json:"open"`          // 시가
	High         decimal.Decimal `gorm:"column:high;type:decimal(11,3)" json:"high"`          // 고가
	Low          decimal.Decimal `gorm:"column:low;type:decimal(11,3)" json:"low"`            // 저가
	Close        decimal.Decimal `gorm:"column:close;type:decimal(11,3)" json:"close"`        // 종가(현재가)
	Percent      string          `gorm:"column:percent;type:float;not null" json:"percent"`   // 등락률
	Volume       string          `gorm:"column:volume;type:bigint;not null" json:"volume"`    // 거래량
}

type OneTickerResponse struct {
	Symbol         string                   `gorm:"column:symbol;type:varchar(12);not null;primary_key" json:"symbol"` //index
	Merket         int                      `gorm:"column:market;type:tinyint;not null" json:"market"`                 // 0 : 코스피 , 1 : 코스닥
	Name           string                   `gorm:"column:name;type:varchar(100);not null" json:"name"`                // 종목명
	Bps            decimal.Decimal          `gorm:"column:bps;type:decimal(11,2);default:0" json:"bps"`
	Per            decimal.Decimal          `gorm:"column:per;type:decimal(11,2);default:0" json:"per"` // 표시
	Pbr            decimal.Decimal          `gorm:"column:pbr;type:decimal(11,2);default:0" json:"pbr"` // 표시
	Eps            decimal.Decimal          `gorm:"column:eps;type:decimal(11,2);default:0" json:"eps"`
	Div            decimal.Decimal          `gorm:"column:div;type:decimal(11,2);default:0" json:"div"`
	Dps            decimal.Decimal          `gorm:"column:dps;type:decimal(11,2);default:0" json:"dps"`
	Sector         string                   `gorm:"column:sector;type:varchar(150)" json:"sector"` // 표시
	Industry       string                   `gorm:"column:industry;type:varchar(150)" json:"industry"`
	ListingDate    time.Time                `gorm:"column:listing_date;default:null" json:"listing_date"`                // 상장일
	SettleMonth    string                   `gorm:"column:settle_month;type:varchar(10)" json:"settle_month"`            // 결산월
	Representative string                   `gorm:"column:representative;type:varchar(100)" json:"representative"`       // 대표자 (표시할 필요없습니다.)
	Homepage       string                   `gorm:"column:homepage;type:varchar(200)" json:"homepage"`                   // 홈페이지 (로고 url)
	MarketCap      decimal.Decimal          `gorm:"column:market_cap;type:decimal(17,0);default:null" json:"market_cap"` // 시가총액
	Chart          []OneTickerChartResponse `gorm:"foreignkey:TickerSymbol" json:"chart"`
}

type SearchTickerResponse struct {
	Symbol   string `gorm:"column:symbol;type:varchar(12);not null;primary_key" json:"symbol"` //코드
	Merket   int    `gorm:"column:market;type:tinyint;not null" json:"market"`                 // 0 코스피 , 1 코스닥
	Name     string `gorm:"column:name;type:varchar(100);not null" json:"name"`                // 종목명
	Homepage string `gorm:"column:homepage;type:varchar(200)" json:"homepage"`                 // 홈페이지 (로고 url)
}
