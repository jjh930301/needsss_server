package routers

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jjh930301/needsss/gin/repositories"
	"github.com/jjh930301/needsss_common/res"
)

// @Tags comment
// @Summary 티커 코멘트 목록
// @Description 2000 성공
// @Description 4101 required login
// @Description 4004 count , created_at error
// @Accept json
// @Produce json
// @Router /ticker/comment/{symbol} [get]
// @Param symbol path string true "symbol"
// @Param count query int true "count"
// @Param created_at query string false "created_at"
// @Success 200 {object} []responses.TickerCommentsResponse
func GetTickerComments(c *gin.Context) {
	symbol := c.Param("symbol")
	count := c.Request.URL.Query().Get("count")
	intCount, err := strconv.Atoi(count)
	if err != nil {
		res.BadRequest(c, "Required count", 4004)
		return
	}
	createdAt := c.Request.URL.Query().Get("created_at")
	//2023-01-15T09:38:18Z
	if createdAt != "" {
		lastTime, err := time.Parse("2006-01-02 15:04:05", createdAt)
		if err != nil {
			res.BadRequest(c, "Time format 2006-01-02 15:04:05", 4004)
			return
		}

		comments := repositories.GetTickerComments(symbol, intCount, lastTime)
		if comments == nil {
			var empty []string
			res.Ok(c, "Empty comments", empty, 2001)
			return
		} else {
			res.Ok(c, "Successfully getting comments", comments, 2000)
			return
		}
	}
	comments := repositories.GetTickerComments(symbol, intCount, nil)

	res.Ok(c, "Successfully getting comments", comments, 2000)

}
