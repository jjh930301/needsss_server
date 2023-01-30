package comment

import (
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jjh930301/needsss_common/res"
	"github.com/jjh930301/needsss_common/structs"
	uuid "github.com/satori/go.uuid"
)

// @Tags comment
// @Summary 티커 코멘트 목록
// @Description 2000 성공
// @Description 4101 required login
// @Description 4004 count , created_at error
// @Accept json
// @Produce json
// @Router /ticker/comment/{ticker} [get]
// @Param ticker path string true "code"
// @Param count query int true "count"
// @Param created_at query string true "created_at"
// @Success 200 {object} []TickerCommentsResponse
func GetTickerComments(c *gin.Context) {
	code := c.Param("ticker")
	count := c.Request.URL.Query().Get("count")
	createdAt := c.Request.URL.Query().Get("created_at")
	//2023-01-15T09:38:18Z
	lastTime, err := time.Parse("2006-01-02T15:04:05Z", createdAt)
	if err != nil {
		fmt.Println(err)
		res.BadRequest(c, "Required Time format 2006-01-02T15:04:05Z", 4004)
		panic(err)
	}
	intCount, err := strconv.Atoi(count)
	if err != nil {
		res.BadRequest(c, "Required count", 4004)
		panic(err)
	}
	comments := getTickerComments(code, lastTime, intCount)
	if comments == nil {
		var empty []string
		res.Ok(c, "Empty comments", empty, 2001)
		panic(nil)
	}

	res.Ok(c, "Successfully getting comments", comments, 2000)

}

// @Tags comment
// @Summary 티커 코멘트 등록
// @Description 2000 성공
// @Description 4101 required login
// @Description 4003 empty comment or code
// @Description 4004 missing bodies
// @Accept json
// @Produce json
// @Security BearerAuth
// @Router /ticker/comment [post]
// @Param data body NewCommentBody true "ticker comment"
// @Success 200 {object} []TickerCommentsResponse
func NewComment(c *gin.Context) {
	user, userErr := c.Keys["user"].(structs.AuthClaim)
	if !userErr {
		res.Forbidden(c, "Required login", 4101)
		panic(userErr)
	}
	var body NewCommentBody
	if err := c.ShouldBindJSON(&body); err != nil {
		res.BadRequest(c, "Missing bodies", 4004)
		panic(err)
	}
	if body.Code == "" || body.Comment == "" {
		res.BadRequest(c, "Empty values", 4003)
		panic(nil)
	}
	uuid, _ := uuid.FromString(user.UserID)
	comments := newTickerComment(uuid, body)

	res.Ok(c, "Successfully new comments", comments, 2000)
}
