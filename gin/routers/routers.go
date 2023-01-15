package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jjh930301/market/common/constants"
	"github.com/jjh930301/market/common/middleware"
	authRouter "github.com/jjh930301/market/routers/auth"
	interestRouter "github.com/jjh930301/market/routers/interest"
	tickerRouter "github.com/jjh930301/market/routers/ticker"
	commentRouter "github.com/jjh930301/market/routers/ticker/comment"
	userRouter "github.com/jjh930301/market/routers/user"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health/check", HealthCheck)
	auth := r.Group(constants.AuthGroup)
	{
		auth.POST(constants.Token, authRouter.Token)
		auth.POST(constants.Regist, authRouter.Register)
		auth.GET(constants.Login, authRouter.Login)
	}

	ticker := r.Group(constants.TickerGroup)
	{
		ticker.GET(constants.GetTicker, tickerRouter.FindOne)
		ticker.GET(constants.GetChart, tickerRouter.GetTickerChart)
		ticker.GET(constants.TickerSearch, tickerRouter.SearchTicker)
		comment := ticker.Group(constants.CommentGroup)
		{
			comment.POST(constants.Default, middleware.Authenticator, commentRouter.NewComment)
			comment.GET(constants.GetTickerComments, commentRouter.GetTickerComments)
		}
	}

	interest := r.Group(constants.InterestGroup)
	{
		interest.GET(constants.Default, interestRouter.GetList)
		interest.POST(constants.Default, middleware.Authenticator, interestRouter.SetList)
		interest.DELETE(constants.Default, middleware.Authenticator, interestRouter.DeleteList)
		interest.PUT(constants.SaleInterestTicker, middleware.Authenticator, interestRouter.SaleInterest)
	}
	user := r.Group(constants.UserGroup)
	{
		user.PUT(constants.SetNickname, middleware.Authenticator, userRouter.SetNickname)
	}

	return r
}
