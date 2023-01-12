package ticker

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jjh930301/market/common/res"
)

// @Tags ticker
// @Summary 종목 가져오기
// @Description 2000 성공
// @Accept json
// @Produce json
// @Param ticker path string true "code"
// @Param offset query int false "offset count"
// @Router /ticker/{ticker} [get]
func FindOne(c *gin.Context) {
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
	var param string = c.Param("ticker")
	var ticker *OneTickerResponse
	var repoErr error
	ticker, repoErr = findOne(param, intOffset)
	if repoErr != nil {
		res.BadRequest(
			c,
			"Cannot found list",
			4002,
		)
		panic(nil)
	}
	res.Ok(
		c,
		"Successfully getting tickers",
		ticker,
		2000,
	)
}
