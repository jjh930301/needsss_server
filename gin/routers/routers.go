package routers

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jjh930301/needsss/gin/endpoint"
	"github.com/jjh930301/needsss/gin/repositories"
	"github.com/jjh930301/needsss_common/middleware"
)

var userRepo repositories.UserRepository

func InitRouter() *gin.Engine {
	r := gin.New()
	if os.Getenv("ENV") == "dev" {
		r.Use(gin.LoggerWithConfig(gin.LoggerConfig{
			SkipPaths: []string{"/health/check"},
		}))
	}

	r.Use(cors.New(middleware.CORSConfig()))

	r.GET("/health/check", HealthCheck)
	auth := r.Group(endpoint.AuthGroup)
	{
		auth.POST(endpoint.Token, Token)
		auth.GET(endpoint.GoogleIDToken, GoogleIDToken)
	}

	ticker := r.Group(endpoint.TickerGroup)
	{
		ticker.GET(endpoint.GetTicker, FindOne)
		ticker.GET(endpoint.GetChart, GetTickerChart)
		ticker.GET(endpoint.TickerSearch, SearchTicker)
		comment := ticker.Group(endpoint.CommentGroup)
		{
			comment.GET(endpoint.GetTickerComments, GetTickerComments)
		}
	}

	interest := r.Group(endpoint.InterestGroup)
	{
		interest.GET(endpoint.Default, GetList)
		interest.POST(endpoint.Default, middleware.Authenticator, SetList)
		interest.DELETE(endpoint.Default, middleware.Authenticator, DeleteList)
		interest.PUT(endpoint.SaleInterestTicker, middleware.Authenticator, SaleInterest)
	}
	user := r.Group(endpoint.UserGroup)
	{
		user.PUT(endpoint.SetNickname, middleware.Authenticator, SetNickname)
	}
	holidays := r.Group(endpoint.HolidayGroup)
	{
		holidays.POST(endpoint.Default, middleware.Authenticator, InsertHoliday)
	}
	return r
}
