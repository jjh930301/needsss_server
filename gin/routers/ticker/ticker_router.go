package ticker

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jjh930301/market/routers/logs"
	"github.com/jjh930301/market/vars"
	"github.com/jjh930301/needsss_common/res"
)

// @Tags ticker
// @Summary 종목 가져오기
// @Description 2000 성공
// @Description 4001 required count
// @Description 4002 cannot found list
// @Accept json
// @Produce json
// @Param ticker path string true "code"
// @Param count query int false "count"
// @Router /ticker/{ticker} [get]
// @Success 200 {object} OneTickerResponse
func FindOne(c *gin.Context) {
	count := c.Request.URL.Query().Get("count")
	intCount, err := strconv.Atoi(count)
	if err != nil {
		res.BadRequest(
			c,
			"Required count",
			4001,
		)
	}

	var param string = c.Param("ticker")
	var ticker *OneTickerResponse
	var repoErr error
	ticker, repoErr = findOne(param, intCount)
	if repoErr != nil {
		res.BadRequest(
			c,
			"Cannot found list",
			4002,
		)
		panic(nil)
	}
	go logs.Log(c.ClientIP(), vars.GET, vars.TickerGroup+vars.GetTicker)
	res.Ok(
		c,
		"Successfully getting tickers",
		ticker,
		2000,
	)
}

// @Tags ticker
// @Summary 종목 차트 가져오기
// @Description 2000 성공
// @Description 2001 empty response
// @Description 4001 missing count
// @Description 4002 date format error yyyy-mm-dd
// @Accept json
// @Produce json
// @Param ticker path string true "code"
// @Param date query string false "마지막 날짜"
// @Param count query int false "count"
// @Router /ticker/chart/{ticker} [get]
// @Success 200 {object} []OneTickerChartResponse
func GetTickerChart(c *gin.Context) {
	strDate := c.Request.URL.Query().Get("date")
	count := c.Request.URL.Query().Get("count")
	intCount, err := strconv.Atoi(count)
	if err != nil {
		res.BadRequest(
			c,
			"Required count",
			4001,
		)
		panic(err)
	}
	date, err := time.Parse("2006-01-02", strDate)
	if err != nil {
		res.BadRequest(
			c,
			"Out of format",
			4002,
		)
		panic(err)
	}

	var param string = c.Param("ticker")
	chart := findChart(param, intCount, date)
	if chart == nil {
		var empty []string
		res.Ok(c, "Empty chart", gin.H{"chart": empty}, 2001)
		panic(nil)
	}
	go logs.Log(c.ClientIP(), vars.GET, vars.TickerGroup+vars.GetChart)
	res.Ok(c, "Successfully getting charts", chart, 2000)
}

// @Tags ticker
// @Summary 종목 검색
// @Description 2000 성공
// @Description 4001 offset and count are int type
// @Accept json
// @Produce json
// @Param word query string true "검색할 종목명 , 코드"
// @Param offset query int false "가지고 있는 list.length"
// @Param count query int false "count"
// @Router /ticker/search [get]
// @Success 200 {object} []SearchTickerResponse
func SearchTicker(c *gin.Context) {
	word := c.Request.URL.Query().Get("word")
	offset := c.Request.URL.Query().Get("offset")
	count := c.Request.URL.Query().Get("count")
	intOffset, err := strconv.Atoi(offset)
	if err != nil {
		res.BadRequest(c, "offset and count are int type", 4001)
		panic(err)
	}
	intCount, err := strconv.Atoi(count)
	if err != nil {
		res.BadRequest(c, "offset and count are int type", 4001)
		panic(err)
	}
	tickers := searchTicker(word, intOffset, intCount)

	res.Ok(c, "Successfully getting tickers", tickers, 2000)

}
