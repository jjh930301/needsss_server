package interest

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jjh930301/market/common/constants"
	"github.com/jjh930301/market/common/res"
	"github.com/jjh930301/market/common/structs"
	"github.com/jjh930301/market/routers/logs"
)

// @Tags interest
// @Summary 관심종목
// @Description 2000 성공
// @Accept json
// @Produce json
// @Param offset query int false "offset count"
// @Router /interest [get]
// @Success 200 {object} InterestListResponse
func GetList(c *gin.Context) {
	offset := c.Request.URL.Query().Get("offset")
	intOffset, err := strconv.Atoi(offset)
	if err != nil {
		res.BadRequest(
			c,
			"Required offset",
			4001,
		)
		panic(nil)
	}
	interestes, err := getList(intOffset)
	if err != nil {
		res.ServerError(c)
		panic(err)
	}
	go logs.Log(c.ClientIP(), constants.GET, constants.InterestGroup+constants.Default)
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
// @Param data body SetInterestBody true "add interest".
// @Router /interest [post]
// @Success 200 {object} InterestListResponse
// @Security BearerAuth
func SetList(c *gin.Context) {
	user, userErr := c.Keys["user"].(structs.AuthClaim)
	if !userErr {
		res.Forbidden(c, "Required login", 4101)
		panic(userErr)
	}
	var body SetInterestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		res.BadRequest(c, "Missing Bodies", 4001)
		panic(nil)
	}
	// var chart *models.KrTickerChartsModel
	_, err := setList(user.UserID, body.Code)
	if err != nil {
		res.BadRequest(c, "Cannot add ticker", 4002)
	}
	go logs.Log(c.ClientIP(), constants.POST, constants.InterestGroup+constants.Default)
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
// @Param data body DeleteIntereestBody true "delete interest".
// @Router /interest [delete]
// @Security BearerAuth
func DeleteList(c *gin.Context) {
	user, userErr := c.Keys["user"].(structs.AuthClaim)
	if !userErr {
		res.Forbidden(c, "Required login", 4101)
		panic(userErr)
	}

	var body DeleteIntereestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		res.BadRequest(c, "Missing bodies", 4001)
		panic(err)
	}
	err := deleteList(user.UserID, &body)
	if err != nil {
		res.BadRequest(c, "Cannot delete ticker", 4002)
		panic(err)
	}
	go logs.Log(c.ClientIP(), constants.DELETE, constants.InterestGroup+constants.Default)
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
// @Param data body SaleInterestBody true "sale interest".
// @Router /interest/sale [put]
// @Security BearerAuth
func SaleInterest(c *gin.Context) {
	user, userErr := c.Keys["user"].(structs.AuthClaim)
	if !userErr {
		res.Forbidden(c, "Required login", 4101)
		panic(userErr)
	}
	var body SaleInterestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		res.BadRequest(c, "Missing bodies", 4002)
		panic(err)
	}
	krDate := time.Now().In(constants.KrTime()).Format("2006-01-02")
	endTime, parseErr := time.Parse("2006-01-02T15:04:05.000Z", krDate+"T15:20:00.000Z")
	if parseErr != nil {
		res.BadRequest(c, "test", 4444)
		panic(parseErr)
	}
	now := time.Now().In(constants.KrTime())
	if endTime.Before(now) {
		res.BadRequest(c, "The time to sale has over", 4005)
		panic(nil)
	}

	err := saleInterest(user.UserID, body, now)
	if !err {
		res.BadRequest(c, "Cannot sale my interest", 4004)
		panic(err)
	}
	go logs.Log(c.ClientIP(), constants.PUT, constants.InterestGroup+constants.SaleInterestTicker)
	res.Ok(c, "Successfully sale my interest", gin.H{"result": true}, 2000)
}
