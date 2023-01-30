package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/jjh930301/needsss/socket/handlers"
	"github.com/jjh930301/needsss_common/database"
	"github.com/jjh930301/needsss_common/middleware"
	"github.com/jjh930301/needsss_common/utils"
)

func main() {
	utils.OauthInit()
	database.InitDb()
	r := gin.Default()

	r.Use(cors.New(middleware.CORSConfig()))
	// Use Transport Pooling, because Websocket Transport still Error
	server := handlers.Server()
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("Connected:", s.ID())
		return nil
	})
	server.OnEvent("/", "comment", handlers.HandleComment)
	server.OnEvent("/", "ticker", handlers.HandleTickerRoom)
	server.OnDisconnect("/", func(s socketio.Conn, msg string) {
		fmt.Println("DisConneceted", s.ID())
	})
	go server.Serve()
	defer server.Close()

	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))
	// router.POST("/socket.io/", func(context *gin.Context) {
	// 	server.ServeHTTP(context.Writer, context.Request)
	// })
	// Run Gin
	r.Run(":7007")
}
