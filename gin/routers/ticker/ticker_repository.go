package ticker

import (
	"github.com/jjh930301/market/common/database"
	"github.com/jjh930301/market/common/models"
	"gorm.io/gorm"
)

func findOne(code string, offset int) (*OneTickerResponse, error) {
	var err error
	var ticker *OneTickerResponse
	var charts *oneTickerChartModel
	var count int = 120
	if code == "" {
		return nil, err
	}

	// var result *FirstTickerModel
	database.DB.Model(&models.KrTickerModel{}).Preload("Chart", func(chart *gorm.DB) *gorm.DB {
		return chart.Model(&models.KrTickerChartsModel{}).Order(
			"date desc",
		).Limit(count).Offset(offset).Find(&charts)
	}).Where(&models.KrTickerModel{Symbol: code}).First(&ticker)
	return ticker, nil
}
