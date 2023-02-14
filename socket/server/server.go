package server

import (
	"net/http"

	socketio "github.com/googollee/go-socket.io"
	"github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"
	"github.com/googollee/go-socket.io/engineio/transport/polling"
	"github.com/googollee/go-socket.io/engineio/transport/websocket"
)

var Socket *socketio.Server

func Server() *socketio.Server {
	pollingDefault := polling.Default

	wsDefault := websocket.Default
	wsDefault.CheckOrigin = func(req *http.Request) bool {
		return true
	}
	Socket = socketio.NewServer(
		&engineio.Options{
			Transports: []transport.Transport{
				pollingDefault,
				wsDefault,
			},
		},
	)

	return Socket
}
