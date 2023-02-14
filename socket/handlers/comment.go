package handlers

import (
	"encoding/json"
	"fmt"

	socketio "github.com/googollee/go-socket.io"
	"github.com/jjh930301/needsss/socket/jsons"
	"github.com/jjh930301/needsss/socket/repositories"
	"github.com/jjh930301/needsss/socket/server"
)

func HandleComment(s socketio.Conn, msg string) {
	var comment jsons.CommentJson
	err := json.Unmarshal([]byte(msg), &comment)
	if err != nil {
		return
	}
	if comment.UserId == "" {
		return
	}
	newComment := repositories.NewTickerComment(comment)
	if newComment == nil {
		return
	}
	new, _ := json.Marshal(newComment)

	server.Socket.ForEach("/", comment.Symbol, func(c socketio.Conn) {
		fmt.Println("id:::", c.ID())
		go c.Emit(comment.Symbol, string(new))
	})
}
