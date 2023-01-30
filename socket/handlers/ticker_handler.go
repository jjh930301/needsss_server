package handlers

import (
	"encoding/json"

	socketio "github.com/googollee/go-socket.io"
)

func HandleTickerRoom(s socketio.Conn, msg string) string {
	json, err := json.Marshal(msg)
	if err != nil {
		return ""
	}
	s.SetContext(json)
	return msg
}
