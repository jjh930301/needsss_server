package repositories

import (
	"github.com/jjh930301/needsss/socket/dao"
	"github.com/jjh930301/needsss_common/database"
	"github.com/jjh930301/needsss_common/models"
	"gorm.io/gorm"
)

func FindInterest(symbol string) *[]dao.InterestResponse {
	var interestes []dao.Interest
	var userModel dao.InterestUserModel
	var tickerModel dao.InterestTickerModel
	var response []dao.InterestResponse

	database.DB.Model(&models.InterestedTickerModel{}).Preload("User", func(user *gorm.DB) *gorm.DB {
		return user.Model(&models.UserModel{}).Find(&userModel)
	}).Preload("Ticker", func(ticker *gorm.DB) *gorm.DB {
		return ticker.Model(&models.KrTickerModel{}).Find(&tickerModel)
	}).Where(&models.InterestedTickerModel{
		TickerSymbol: symbol,
	}).Find(&interestes)
	if len(interestes) == 0 {
		return nil
	}
	var recentOne dao.InterestTickerChartModel
	database.DB.Model(&models.KrTickerChartsModel{}).Where(&models.KrTickerChartsModel{
		Symbol: symbol,
	}).Order("date desc").First(&recentOne)

	for _, i := range interestes {
		var chartModel dao.InterestTickerChartModel
		database.DB.Model(&models.KrTickerChartsModel{}).Where(&models.KrTickerChartsModel{
			Symbol: i.TickerSymbol,
		}).Order("date desc").First(&chartModel)
		response = append(response, dao.InterestResponse{
			Ticker: i,
			Recent: chartModel,
		})
	}
	return &response
}
