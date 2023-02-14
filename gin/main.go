package main

import (
	"os"

	"github.com/jjh930301/needsss/gin/docs"
	routers "github.com/jjh930301/needsss/gin/routers"
	"github.com/jjh930301/needsss_common/database"
	"github.com/jjh930301/needsss_common/models"
	"github.com/jjh930301/needsss_common/utils"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// _ "github.com/swaggo/gin-swagger/example/basic/docs"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	utils.OauthInit()
	database.InitDb()
	if os.Getenv("ENV") != "production" {
		database.DB.AutoMigrate(
			&models.HolidayModel{},
		)
	}
	r := routers.InitRouter()

	docs.SwaggerInfo.Title = "관심종목 Api Documentation"
	docs.SwaggerInfo.Description = `
		Use Bearer Token
		market type
		0 = KOSPI
		1 = KOSDAQ
	`

	// 127.0.0.1:7070/docs/index.html
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":7070")
}
