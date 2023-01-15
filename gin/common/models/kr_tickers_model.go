package models

import (
	"time"

	"github.com/jjh930301/market/common/constants"
	"github.com/shopspring/decimal"
)

type KrTickerModel struct {
	Symbol         string                  `gorm:"column:symbol;type:varchar(12);not null;primary_key" json:"symbol,omitempty"` //index
	Merket         int                     `gorm:"column:market;type:tinyint" json:"market,omitempty"`
	Name           string                  `gorm:"column:name;type:varchar(100);not null" json:"name,omitempty"` //index
	Bps            decimal.Decimal         `gorm:"column:bps;type:decimal(11,2);default:0" json:"bps,omitempty"`
	Per            decimal.Decimal         `gorm:"column:per;type:decimal(11,2);default:0" json:"per,omitempty"`
	Pbr            decimal.Decimal         `gorm:"column:pbr;type:decimal(11,2);default:0" json:"pbr,omitempty"`
	Eps            decimal.Decimal         `gorm:"column:eps;type:decimal(11,2);default:0" json:"eps,omitempty"`
	Div            decimal.Decimal         `gorm:"column:div;type:decimal(11,2);default:0" json:"div,omitempty"`
	Dps            decimal.Decimal         `gorm:"column:dps;type:decimal(11,2);default:0" json:"dps,omitempty"`
	Sector         string                  `gorm:"column:sector;type:varchar(150)" json:"sector,omitempty"`
	Industry       string                  `gorm:"column:industry;type:varchar(150)" json:"industry,omitempty"`
	ListingDate    time.Time               `gorm:"column:listing_date;default:null" json:"listing_date,omitempty"`
	SettleMonth    string                  `gorm:"column:settle_month;type:varchar(10)" json:"settle_month,omitempty"`
	Representative string                  `gorm:"column:representative;type:varchar(100)" json:"representative,omitempty"`
	Homepage       string                  `gorm:"column:homepage;type:varchar(200)" json:"homepage,omitempty"`
	MarketCap      decimal.Decimal         `gorm:"column:market_cap;type:decimal(17,0);default:null" json:"market_cap,omitempty"`
	Region         string                  `gorm:"column:region;type:varchar(10)" json:"region,omitempty"`
	Chart          []KrTickerChartsModel   `gorm:"foreignKey:TickerSymbol" json:"chart,omitempty"`
	Comments       []KrTickerCommentModel  `gorm:"foreignKey:TickerSymbol" json:"comments,omitempty"`
	Interestes     []InterestedTickerModel `gorm:"foreignKey:TickerSymbol" json:"interests,omitempty"`
}

func (KrTickerModel) TableName() string {
	return constants.KrTickerTableName
}
