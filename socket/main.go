package main

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"github.com/jjh930301/needsss/socket/handlers"
	"github.com/jjh930301/needsss/socket/server"
	"github.com/jjh930301/needsss_common/database"
	"github.com/jjh930301/needsss_common/middleware"
	"github.com/jjh930301/needsss_common/utils"
)

func main() {
	utils.OauthInit()
	var user string
	if os.Getenv("MYSQL_USER") == "" {
		user = "root"
	} else {
		user = os.Getenv("MYSQL_USER")
	}
	database.InitDb(
		user,
		os.Getenv("MYSQL_ROOT_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_DATABASE"),
	)
	r := gin.Default()

	r.Use(cors.New(middleware.CORSConfig()))
	server := server.Server()
	server.OnConnect("/", func(s socketio.Conn) error {
		s.Join("main")
		return nil
	})
	server.OnEvent("/", "ticker", handlers.JoinTickerRoom)
	server.OnEvent("/", "leave", handlers.LeaveTickerRoom)
	server.OnEvent("/", "comment", handlers.HandleComment)
	server.OnEvent("/", "interest", handlers.HandleRealTime)

	server.OnDisconnect("/", func(s socketio.Conn, msg string) {
		s.LeaveAll()
	})

	go server.Serve()
	defer server.Close()

	r.GET("/socket.io/*any", gin.WrapH(server))
	r.POST("/socket.io/*any", gin.WrapH(server))
	// Run Gin
	r.Run(":7007")
}
