package interest

import (
	"fmt"
	"time"

	"github.com/jjh930301/market/common/constants"
	"github.com/jjh930301/market/common/database"
	"github.com/jjh930301/market/common/models"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func getList(offset int) (*[]InterestListResponse, error) {
	var err error
	var interestes []interestList
	var userModel interestUserModel
	var responseModel []InterestListResponse
	uuid, err := uuid.FromString("b733cfae-1310-43ba-8dd1-ee337a64b980")
	if err != nil {
		return nil, err
	}
	database.DB.Model(&models.InterestedTickerModel{}).Preload("User", func(user *gorm.DB) *gorm.DB {
		return user.Model(&models.UserModel{}).Where(&models.UserModel{
			ID: uuid,
		}).First(&userModel)
	}).Where(
		&models.InterestedTickerModel{UserId: uuid},
	).Order("date_time desc").Limit(20).Offset(offset).Find(&interestes)
	for _, i := range interestes {
		var chartModel interestTickerChartModel
		database.DB.Model(&models.KrTickerChartsModel{}).Where(&models.KrTickerChartsModel{
			Symbol: i.TickerSymbol,
		}).Order("date desc").First(&chartModel)
		responseModel = append(responseModel, InterestListResponse{
			Ticker: i,
			Recent: chartModel,
		})
	}
	return &responseModel, nil
}

func setList(userId string, code string) (bool, error) {
	uuid, _ := uuid.FromString(userId)
	var ticker *models.KrTickerModel
	var chart *models.KrTickerChartsModel
	krTime := time.Now().In(constants.KrTime())
	database.DB.Model(&models.KrTickerChartsModel{}).Preload(
		"Ticker",
		func(model *gorm.DB) *gorm.DB {
			return model.Model(&models.KrTickerModel{}).Where(
				&models.KrTickerModel{Symbol: code},
			).First(&ticker)
		},
	).Where(
		&models.KrTickerChartsModel{Symbol: code},
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
func deleteList(id string, body *DeleteIntereestBody) error {

	uuid, _ := uuid.FromString(id)
	date, parseErr := time.Parse("2006-01-02", body.Date)
	if parseErr != nil {
		return fmt.Errorf("Cannot string to date")
	}
	err := database.DB.Where(&models.InterestedTickerModel{
		UserId:       uuid,
		TickerSymbol: body.Code,
		Date:         date,
	}).Delete(&models.InterestedTickerModel{}).Error
	fmt.Println(err)
	return err
}
