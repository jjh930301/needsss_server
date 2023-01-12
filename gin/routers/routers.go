package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/jjh930301/market/common/middleware"
	authRouter "github.com/jjh930301/market/routers/auth"
	interestRouter "github.com/jjh930301/market/routers/interest"
	tickerRouter "github.com/jjh930301/market/routers/ticker"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/health/check", HealthCheck)
	auth := r.Group("/auth")
	{
		auth.POST("/regist", authRouter.Register)
		auth.GET("/login", authRouter.Login)
	}

	ticker := r.Group("/ticker")
	{
		ticker.GET("/:ticker", tickerRouter.FindOne)
	}

	interest := r.Group("interest")
	{
		interest.GET("", interestRouter.GetList)
		interest.POST("", middleware.Authenticator, interestRouter.SetList)
	}

	return r
}
