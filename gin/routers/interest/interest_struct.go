package interest

import (
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/shopspring/decimal"
)

type interestUserModel struct {
	ID       uuid.UUID `gorm:"column:id;type:varchar(36);primary_key" json:"-"`
	NickName string    `gorm:"column:nickname;type:varchar(30);default:null" json:"nickname"` // 등록한 사용자의 닉네임
}

type interestTickerModel struct {
	Symbol   string `gorm:"column:symbol;type:varchar(12);not null;primary_key" json:"-"`
	Homepage string `gorm:"column:homepage;type:varchar(200)" json:"homepage"` // 홈페이지
}

type interestTickerChartModel struct {
	TickerSymbol string          `gorm:"column:ticker_symbol;type:varchar(12);not null" json:"-"`
	Close        decimal.Decimal `gorm:"column:close;type:decimal(11,3)" json:"close"`      // 현재가
	Percent      string          `gorm:"column:percent;type:float;not null" json:"percent"` // 현재 등락률
	Volume       string          `gorm:"column:volume;type:bigint;not null" json:"volume"`  // 현재 거래량
}

type interestList struct {
	UserId       uuid.UUID           `gorm:"column:user_id;type:varchar(36)" json:"-"`
	User         interestUserModel   `json:"user"`
	TickerSymbol string              `gorm:"column:symbol;type:varchar(12);not null;primary_key" json:"symbol"` //종목코드
	Ticker       interestTickerModel `json:"ticker"`
	DateTime     time.Time           `json:"date_time"` // 등록된 시점
	Market       int                 `json:"market"`    // 0 KOSPI 1 KOSDAQ
	Type         int                 `json:"type"`      // 현재 사용하지 않습니다.
	Name         string              `json:"name"`      // 종목이름
	Close        decimal.Decimal     `json:"close"`     // 등록한 시점의 가격
	Percent      string              `json:"percent"`   // 등록한 시점의 등락률
	Volume       string              `json:"volume"`    // 등록된 시점의 거래량
}

type InterestListResponse struct {
	Ticker interestList             `json:"ticker"`
	Recent interestTickerChartModel `json:"recent"`
}

type SetInterestBody struct {
	Code string `json:"code"`
}

type DeleteIntereestBody struct {
	Code string `json:"code"`
	Date string `json:"date"`
}

type recentChart struct {
	Close decimal.Decimal `gorm:"column:close;type:decimal(11,3)" json:"close"`
}

type SaleInterestBody struct {
	Code string `json:"code"`
	Date string `json:"date"`
}
