package handlers

import (
	"encoding/json"

	socketio "github.com/googollee/go-socket.io"
	"github.com/jjh930301/needsss/socket/repositories"
	"github.com/jjh930301/needsss/socket/server"
)

func HandleRealTime(s socketio.Conn, symbol string) {
	interest := repositories.FindInterest(symbol)
	if interest == nil {
		return
	}
	data, _ := json.Marshal(interest)
	server.Socket.ForEach("/", "interest", func(c socketio.Conn) {
		go c.Emit("interest", string(data))
	})
}
