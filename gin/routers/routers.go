package routers

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	authRouter "github.com/jjh930301/market/routers/auth"
	interestRouter "github.com/jjh930301/market/routers/interest"
	tickerRouter "github.com/jjh930301/market/routers/ticker"
	commentRouter "github.com/jjh930301/market/routers/ticker/comment"
	userRouter "github.com/jjh930301/market/routers/user"
	"github.com/jjh930301/market/vars"
	"github.com/jjh930301/needsss_common/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.New(middleware.CORSConfig()))

	r.GET("/health/check", HealthCheck)
	auth := r.Group(vars.AuthGroup)
	{
		auth.POST(vars.Token, authRouter.Token)
		auth.GET(vars.GoogleIDToken, authRouter.GoogleIDToken)
	}

	ticker := r.Group(vars.TickerGroup)
	{
		ticker.GET(vars.GetTicker, tickerRouter.FindOne)
		ticker.GET(vars.GetChart, tickerRouter.GetTickerChart)
		ticker.GET(vars.TickerSearch, tickerRouter.SearchTicker)
		comment := ticker.Group(vars.CommentGroup)
		{
			comment.POST(vars.Default, middleware.Authenticator, commentRouter.NewComment)
			comment.GET(vars.GetTickerComments, commentRouter.GetTickerComments)
		}
	}

	interest := r.Group(vars.InterestGroup)
	{
		interest.GET(vars.Default, interestRouter.GetList)
		interest.POST(vars.Default, middleware.Authenticator, interestRouter.SetList)
		interest.DELETE(vars.Default, middleware.Authenticator, interestRouter.DeleteList)
		interest.PUT(vars.SaleInterestTicker, middleware.Authenticator, interestRouter.SaleInterest)
	}
	user := r.Group(vars.UserGroup)
	{
		user.PUT(vars.SetNickname, middleware.Authenticator, userRouter.SetNickname)
	}
	return r
}
