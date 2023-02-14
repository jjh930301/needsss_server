package routers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jjh930301/needsss/gin/dto"
	"github.com/jjh930301/needsss/gin/repositories"
	"github.com/jjh930301/needsss_common/res"
)

// @Tags holiday
// @Summary 공휴일 or 매매 늦게 시작하는 날 설정
// @Description 2000 성공
// @Description 4003 missing bodies
// @Description 4004 이미 생성된 공휴일
// @Accept json
// @Produce json
// @Router /holiday [post]
// @Param data body dto.HolidayDto true "holiday"
// @Security BearerAuth
func InsertHoliday(c *gin.Context) {
	var body dto.HolidayDto
	if err := c.ShouldBindJSON(&body); err != nil {
		fmt.Println(err)
		res.BadRequest(c, "missing bodies", 4004)
		return
	}
	date, err := time.Parse("2006-01-02", body.Date)
	if err != nil {
		res.BadRequest(c, "date parse error", 4002)
		return
	}
	fmt.Println(body.OpenedAt)
	body.TDate = date
	if body.OpenedAt == "" {
		openedAt, _ := time.Parse("2006-01-02 15:04:05", body.Date+" 00:00:00")
		body.TOpendAt = openedAt
	} else {
		openedAt, _ := time.Parse("2006-01-02 15:04:05", body.OpenedAt)
		body.TOpendAt = openedAt
	}
	createErr := repositories.CreateHoliday(body)
	if createErr != nil {
		fmt.Println(createErr)
		res.BadRequest(c, "Cannot create new holiday", 4003)
		return
	}
	res.Ok(c, "Successfully create new holiday", gin.H{"result": true}, 2001)
}
