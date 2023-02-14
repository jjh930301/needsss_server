package repositories

import (
	"time"

	"github.com/jjh930301/needsss/gin/responses"
	"github.com/jjh930301/needsss_common/database"
	"github.com/jjh930301/needsss_common/models"
	"gorm.io/gorm"
)

func FindOnebySymbol(
	symbol string,
	count int,
) *responses.OneTickerResponse {
	var ticker responses.OneTickerResponse
	if symbol == "" {
		return nil
	}
	// var result *FirstTickerModel
	database.DB.Model(&models.KrTickerModel{}).Preload("Chart", func(chart *gorm.DB) *gorm.DB {
		var charts responses.OneTickerChartResponse
		return chart.Model(&models.KrTickerChartsModel{}).Order(
			"date desc",
		).Limit(count).Find(&charts)
	}).Where(&models.KrTickerModel{Symbol: symbol}).First(&ticker)
	return &ticker
}

func FindChart(
	symbol string,
	count int,
	date time.Time,
) *[]responses.OneTickerChartResponse {
	var chart *[]responses.OneTickerChartResponse
	query := database.DB.Model(&models.KrTickerChartsModel{}).Where(
		&models.KrTickerChartsModel{Symbol: symbol},
	).Order("date desc")
	if !date.IsZero() {
		query.Where("date < ?", date)
	}
	query.Limit(count).Find(&chart)
	return chart
}

func SearchTicker(
	word string,
	offset int,
	count int,
) *[]responses.SearchTickerResponse {
	var tickers []responses.SearchTickerResponse
	database.DB.Model(&models.KrTickerModel{}).Where(
		"CONCAT(symbol,name) LIKE ?", "%"+word+"%",
	).Order(
		"market_cap desc",
	).Limit(count).Offset(offset).Find(&tickers)

	return &tickers
}
