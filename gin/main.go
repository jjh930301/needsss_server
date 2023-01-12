package main

import (
	"github.com/jjh930301/market/common/database"
	"github.com/jjh930301/market/docs"
	routers "github.com/jjh930301/market/routers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	// _ "github.com/swaggo/gin-swagger/example/basic/docs"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	database.InitDb()
	r := routers.InitRouter()
	docs.SwaggerInfo.Title = "관심종목 Api Documentation"
	docs.SwaggerInfo.Description = `Use Bearer Token`

	// 127.0.0.1:8090/docs/index.html
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8090")
}
