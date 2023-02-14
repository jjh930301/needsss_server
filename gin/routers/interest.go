package routers

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jjh930301/needsss/gin/dto"
	"github.com/jjh930301/needsss/gin/endpoint"
	"github.com/jjh930301/needsss/gin/repositories"
	"github.com/jjh930301/needsss_common/constants"
	"github.com/jjh930301/needsss_common/res"
	"github.com/jjh930301/needsss_common/structs"
)

// @Tags interest
// @Summary 관심종목
// @Description 2000 성공
// @Accept json
// @Produce json
// @Param offset query int false "offset count"
// @Router /interest [get]
// @Success 200 {object} []responses.InterestListResponse
func GetList(c *gin.Context) {
	offset := c.Request.URL.Query().Get("offset")
	intOffset, err := strconv.Atoi(offset)
	if err != nil {
		res.BadRequest(
			c,
			"Required offset",
			4001,
		)
		return
	}
	interestes := repositories.GetList(intOffset)
	go repositories.Log(c.ClientIP(), endpoint.GET, endpoint.InterestGroup+endpoint.Default)
	res.Ok(
		c,
		"Successfully getting list",
		interestes,
		2000,
	)
}

// @Tags interest
// @Summary 관심종목 추가
// @Description 2000 성공
// @Description 4001 missing bodies
// @Description 4002 cannot add ticker
// @Description 4101 required login
// @Accept json
// @Produce json
// @Param data body dto.SetInterestDto true "add interest".
// @Router /interest [post]
// @Success 200 {object} responses.InterestListResponse
// @Security BearerAuth
func SetList(c *gin.Context) {
	user, userErr := c.Keys["user"].(structs.AuthClaim)
	if !userErr {
		res.Forbidden(c, "Required login", 4101)
		return
	}
	var body dto.SetInterestDto
	if err := c.ShouldBindJSON(&body); err != nil {
		res.BadRequest(c, "Missing Bodies", 4001)
		return
	}
	// var chart *models.KrTickerChartsModel
	_, err := repositories.SetList(user.UserID, body.Symbol)
	if err != nil {
		res.BadRequest(c, "Cannot add ticker", 4002)
	}
	go repositories.Log(c.ClientIP(), endpoint.POST, endpoint.InterestGroup+endpoint.Default)
	res.Ok(c, "Succesfully add ticker", gin.H{"result": true}, 2000)
}

// @Tags interest
// @Summary 관심종목 삭제
// @Description 2000 성공
// @Description 4001 missing bodies
// @Description 4002 cannot delete[sql]
// @Description 4101 required login
// @Accept json
// @Produce json
// @Param data body dto.DeleteInterestDto true "delete interest".
// @Router /interest [delete]
// @Security BearerAuth
func DeleteList(c *gin.Context) {
	user, userErr := c.Keys["user"].(structs.AuthClaim)
	if !userErr {
		res.Forbidden(c, "Required login", 4101)
		return
	}

	var body dto.DeleteInterestDto
	if err := c.ShouldBindJSON(&body); err != nil {
		res.BadRequest(c, "Missing bodies", 4001)
		return
	}
	err := repositories.DeleteList(user.UserID, &body)
	if err != nil {
		res.BadRequest(c, "Cannot delete ticker", 4002)
		return
	}
	go repositories.Log(c.ClientIP(), endpoint.DELETE, endpoint.InterestGroup+endpoint.Default)
	res.Ok(c, "Successfully delete ticker", gin.H{"result": true}, 2000)
}

// @Tags interest
// @Summary 관심종목 매도
// @Description 2000 성공
// @Description 4002 missing bodies
// @Description 4101 required login
// @Description 4004 Cannot sale
// @Description 4005 The time to sale has over
// @Accept json
// @Produce json
// @Param data body dto.SaleInterestDto true "sale interest".
// @Router /interest/sale [put]
// @Security BearerAuth
func SaleInterest(c *gin.Context) {
	user, userErr := c.Keys["user"].(structs.AuthClaim)
	if !userErr {
		res.Forbidden(c, "Required login", 4101)
		return
	}
	var body dto.SaleInterestDto
	if err := c.ShouldBindJSON(&body); err != nil {
		res.BadRequest(c, "Missing bodies", 4002)
		return
	}
	krDate := time.Now().In(constants.KrTime()).Format("2006-01-02")
	endTime, parseErr := time.Parse("2006-01-02T15:04:05.000Z", krDate+"T15:20:00.000Z")
	if parseErr != nil {
		res.BadRequest(c, "test", 4444)
		return
	}
	now := time.Now().In(constants.KrTime())
	fmt.Println(now.Weekday().String())
	weekDay := now.Weekday()
	if weekDay.String() == "Sunday" || weekDay.String() == "Saturday" {
		res.BadRequest(c, "Today is sunady or monday", 4006)
		return
	}
	if endTime.Before(now) {
		res.BadRequest(c, "The time to sale has over", 4005)
		return
	}

	flag := repositories.SaleInterest(user.UserID, body, now)
	if !flag {
		res.BadRequest(c, "Cannot sale my interest", 4004)
		return
	}
	go repositories.Log(c.ClientIP(), endpoint.PUT, endpoint.InterestGroup+endpoint.SaleInterestTicker)
	res.Ok(c, "Successfully sale my interest", gin.H{"result": true}, 2000)
}
