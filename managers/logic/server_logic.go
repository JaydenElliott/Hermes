package logic

import (
	"arcstack/arcstack-chat-server/pkg/util/connection"
	"fmt"
	"net/http"
	"time"
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

// Create a new websocket server
func NewWsServer() *WsServer {
	return &WsServer{
		broadcast:  make(chan []byte),
		register:   make(chan *User),
		unregister: make(chan *User),
		users:      make(map[*User]bool),
	}
}

// Run the websocket server and listen for register/unregister requests
func (server *WsServer) Run() {
	fmt.Println("WebSocket Server Initialised and Running")
	for {
		select {
		case user := <-server.register:
			server.users[user] = true
		case user := <-server.unregister:
			// if user exists
			if _, ok := server.users[user]; ok {
				delete(server.users, user)
			} else {
				fmt.Println("Unable to unregister user ... user not found")
			}
		case message := <-server.broadcast:
			server.broadcastToUsers(message)
		}
	}
}

func (server *WsServer) broadcastToUsers(message []byte) {
	for user := range server.users {
		user.dataBuffer <- message
	}
}

func (server *WsServer) ServeWs(w http.ResponseWriter, r *http.Request) {
	wsConnection, err := connection.UpgradeHTTPToWS(w, r)
	if err != nil {
		fmt.Println("Error in establishing websocket connection: ", err)
	}

	user := CreateUser("testUser", wsConnection, server)

	go user.CircularRead(1000, 60*time.Second)
	go user.CircularWrite((60*time.Second*9)/10, 10*time.Second)

	server.register <- user
}
