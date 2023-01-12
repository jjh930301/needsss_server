package interest

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jjh930301/market/common/middleware"
	"github.com/jjh930301/market/common/res"
)

// @Tags interest
// @Summary 관심종목
// @Description 2000 성공 \n
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
	}
	res.Ok(
		c,
		"Successfully getting list",
		interestes,
		2000,
	)
}

// @Tags interest
// @Summary 관심종목 추가
// @Description 2000 성공 \n
// @Accept json
// @Produce json
// @Param data body SetInterestBody true "add interest".
// @Router /interest [post]
// @Success 200 {object} InterestListResponse
// @Security BearerAuth
func SetList(c *gin.Context) {

	user, userErr := c.Keys["user"].(middleware.AuthClaim)
	if !userErr {
		res.Forbidden(c, "Required login", 200)
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
	res.Ok(c, "Succesfully add ticker", gin.H{"result": true}, 2000)
}

// @Tags interest
// @Summary 관심종목 삭제
// @Description 2000 성공 \n
// @Accept json
// @Produce json
// @Param data body DeleteIntereestBody true "delete interest".
// @Router /interest [delete]
// @Security BearerAuth
func DeleteList(c *gin.Context) {
	user, userErr := c.Keys["user"].(middleware.AuthClaim)
	if !userErr {
		res.Forbidden(c, "Required login", 200)
	}

	var body DeleteIntereestBody
	if err := c.ShouldBindJSON(&body); err != nil {
		res.BadRequest(c, "Missing bodies", 4001)
		panic(err)
	}
	err := deleteList(user.UserID, &body)
	if err != nil {
		res.BadRequest(c, "Cannot delete ticker", 4002)
	}
	res.Ok(c, "Successfully delete ticker", gin.H{"result": true}, 2000)
}
