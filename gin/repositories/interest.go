package repositories

import (
	"errors"
	"fmt"
	"time"

	"github.com/jjh930301/needsss/gin/dto"
	"github.com/jjh930301/needsss/gin/responses"
	"github.com/jjh930301/needsss_common/constants"
	"github.com/jjh930301/needsss_common/database"
	"github.com/jjh930301/needsss_common/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func GetList(offset int) *[]responses.InterestListResponse {
	var interestes []dto.InterestList
	var userModel dto.InterestUserModel
	var tickerModel dto.InterestTickerModel
	var responseModel []responses.InterestListResponse
	database.DB.Model(&models.InterestedTickerModel{}).Preload("User", func(user *gorm.DB) *gorm.DB {
		return user.Model(&models.UserModel{}).Find(&userModel)
	}).Preload("Ticker", func(ticker *gorm.DB) *gorm.DB {
		return ticker.Model(&models.KrTickerModel{}).Find(&tickerModel)
	}).Order("date_time desc").Limit(40).Offset(offset).Find(&interestes)

	for _, i := range interestes {
		var chartModel dto.InterestTickerChartModel
		database.DB.Model(&models.KrTickerChartsModel{}).Where(&models.KrTickerChartsModel{
			Symbol: i.TickerSymbol,
		}).Order("date desc").First(&chartModel)
		responseModel = append(responseModel, responses.InterestListResponse{
			Ticker: i,
			Recent: chartModel,
		})
	}
	return &responseModel
}

func SetList(userId string, symbol string) (bool, error) {
	uuid, _ := uuid.FromString(userId)
	var ticker *models.KrTickerModel
	var chart *models.KrTickerChartsModel
	krTime := time.Now().In(constants.KrTime())
	database.DB.Model(&models.KrTickerChartsModel{}).Preload(
		"Ticker",
		func(model *gorm.DB) *gorm.DB {
			return model.Model(&models.KrTickerModel{}).Where(
				&models.KrTickerModel{Symbol: symbol},
			).First(&ticker)
		},
	).Where(
		&models.KrTickerChartsModel{Symbol: symbol},
	).Order(
		"date desc",
	).First(&chart)

	model := &models.InterestedTickerModel{
		UserId:       uuid,
		Date:         chart.Date,
		DateTime:     krTime,
		Type:         0,
		Name:         chart.Ticker.Name,
		Close:        chart.Close,
		Percent:      chart.Percent,
		Volume:       chart.Volume,
		TickerSymbol: chart.Symbol,
	}
	err := database.DB.Create(model).Error
	if err != nil {
		return false, fmt.Errorf("cannot create interest")
	}
	return true, nil
}

// hard delete
func DeleteList(id string, body *dto.DeleteInterestDto) error {
	uuid, _ := uuid.FromString(id)
	date, parseErr := time.Parse("2006-01-02", body.Date)
	if parseErr != nil {
		return errors.New("")
	}
	err := database.DB.Where(&models.InterestedTickerModel{
		UserId:       uuid,
		TickerSymbol: body.Symbol,
		Date:         date,
	}).Delete(&models.InterestedTickerModel{}).Error
	return err
}

func SaleInterest(
	id string,
	body dto.SaleInterestDto,
	now time.Time,
) bool {
	uuid, _ := uuid.FromString(id)
	date, _ := time.Parse("2006-01-02", body.Date)
	var recent dto.RecentChart
	database.DB.Model(&models.KrTickerChartsModel{}).Where(&models.KrTickerChartsModel{
		Symbol: body.Symbol,
	}).Order("date desc").First(&recent)
	result := database.DB.Where(&models.InterestedTickerModel{
		UserId:       uuid,
		Date:         date,
		TickerSymbol: body.Symbol,
	}).Updates(&models.InterestedTickerModel{
		SaledAt:   now,
		SaleClose: recent.Close,
	}).RowsAffected

	return result != 0
}
