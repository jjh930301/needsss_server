package ticker

import (
	"time"

	"github.com/jjh930301/market/common/database"
	"github.com/jjh930301/market/common/models"
	"gorm.io/gorm"
)

func findOne(
	code string,
	count int,
) (*OneTickerResponse, error) {
	var err error
	var ticker *OneTickerResponse
	var charts *OneTickerChartResponse
	if code == "" {
		return nil, err
	}

	// var result *FirstTickerModel
	database.DB.Model(&models.KrTickerModel{}).Preload("Chart", func(chart *gorm.DB) *gorm.DB {
		return chart.Model(&models.KrTickerChartsModel{}).Order(
			"date desc",
		).Limit(count).Find(&charts)
	}).Where(&models.KrTickerModel{Symbol: code}).First(&ticker)
	return ticker, nil
}

func findChart(
	code string,
	count int,
	date time.Time,
) *[]OneTickerChartResponse {
	var chart *[]OneTickerChartResponse
	database.DB.Model(&models.KrTickerChartsModel{}).Where(
		&models.KrTickerChartsModel{Symbol: code},
	).Order("date desc").Where("date < ?", date).Limit(count).Find(&chart)
	return chart
}

func searchTicker(
	word string,
	offset int,
	count int,
) *[]SearchTickerResponse {
	var tickers []SearchTickerResponse
	database.DB.Model(&models.KrTickerModel{}).Where(
		"CONCAT(symbol,name) LIKE ?", "%"+word+"%",
	).Order(
		"market_cap desc",
	).Limit(count).Offset(offset).Find(&tickers)

	return &tickers
}
