package handlers

import (
	"encoding/json"

	socketio "github.com/googollee/go-socket.io"
	"github.com/jjh930301/needsss/socket/jsons"
)

func JoinTickerRoom(s socketio.Conn, msg string) {
	var ticker jsons.TickerJson
	// join ticker chatroom
	err := json.Unmarshal([]byte(msg), &ticker)
	if err != nil {
		return
	}
	s.Join(ticker.Symbol)
	// save user data
	ctx := map[string]string{
		"user_id":   ticker.UserId,
		"fcm_token": ticker.FcmToken,
	}
	s.SetContext(ctx)
}

func LeaveTickerRoom(s socketio.Conn, msg string) {
	var ticker jsons.TickerJson
	// leave ticker chatroom
	err := json.Unmarshal([]byte(msg), &ticker)
	if err != nil {
		return
	}
	s.Leave(ticker.Symbol)
	s.SetContext(nil)
}
