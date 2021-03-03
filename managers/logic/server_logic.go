package logic

import (
	"errors"
	"fmt"
)

type WsServer struct {

	// Registered users (clients)
	users map[*User]bool

	// Incoming user messages
	broadcast chan []byte

	// User register requests
	register chan *User

	// User unregister requests
	unregister chan *User
}

func newWsServer() *WsServer {
	return &WsServer{
		broadcast:  make(chan []byte),
		register:   make(chan *User),
		unregister: make(chan *User),
		users:      make(map[*User]bool),
	}
}

func (server *WsServer) Run() {

}
