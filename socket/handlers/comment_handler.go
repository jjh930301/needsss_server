package handlers

import (
	"fmt"

	socketio "github.com/googollee/go-socket.io"
)

func HandleComment(s socketio.Conn, msg string) string {
	rooms := s.Rooms()
	for _, room := range rooms {
		fmt.Println(room)
	}
	return msg
}
