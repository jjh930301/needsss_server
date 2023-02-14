package repositories

import (
	"github.com/jjh930301/needsss/gin/dto"
	"github.com/jjh930301/needsss_common/database"
	"github.com/jjh930301/needsss_common/models"
)

func CreateHoliday(body dto.HolidayDto) error {
	model := &models.HolidayModel{
		Name:     body.Name,
		Date:     body.TDate,
		OpenedAt: body.TOpendAt,
		Type:     body.Type,
	}
	err := database.DB.Create(model).Error

	return err
}
