package handlers

import (
	"fmt"
	"net/http"
	"time"

	socketio "github.com/googollee/go-socket.io"
	engineio "github.com/googollee/go-socket.io/engineio"
	"github.com/googollee/go-socket.io/engineio/transport"

	engineiopooling "github.com/googollee/go-socket.io/engineio/transport/polling"
)

func Server() *socketio.Server {
	server := socketio.NewServer(&engineio.Options{
		Transports: []transport.Transport{
			&engineiopooling.Transport{
				Client: &http.Client{
					Timeout: time.Hour,
				},
			},
		},
		RequestChecker: checker,
	})
	return server
}

func checker(request *http.Request) (http.Header, error) {
	header := http.Header{}
	fmt.Println("header:::", header)
	header.Set("Access-Control-Allow-Origin", "*")
	header.Set("Access-Control-Allow-Credentials", "false")
	return header, nil
}
