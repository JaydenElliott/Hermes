package logic

import (
	"arcstack/arcstack-chat-server/pkg/setting"
	"arcstack/arcstack-chat-server/pkg/util/connection"
	"fmt"
	"log"
	"net/http"
)

// Websocket server data struct
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

// NewWsServer creates a new websocket server struct and returns it's address.
// Requires no parameters and returns empty channels/maps.
func NewWsServer() *WsServer {
	return &WsServer{
		broadcast:  make(chan []byte),
		register:   make(chan *User),
		unregister: make(chan *User),
		users:      make(map[*User]bool),
	}
}

// Run the websocket server and listen for register/unregister requests.
// Will run continuously.
func (server *WsServer) Run() {
	fmt.Println("WebSocket Server Initialised and Running")
	for {
		select {
		// Register user
		case user := <-server.register:
			server.users[user] = true

		case user := <-server.unregister:
			// Check if user exists
			if _, ok := server.users[user]; ok {
				delete(server.users, user)

			} else {
				// User not found in server
				fmt.Println("Unable to unregister user ... user not found")
			}
		case message := <-server.broadcast:
			server.broadcastToUsers(message)
		}
	}
}

// broadcastToUsers will send the message/messages stored in databuffer to
// all users currently registered on the server.
func (server *WsServer) broadcastToUsers(message []byte) {
	for user := range server.users {
		user.dataBuffer <- message
	}
}

// ServeWs receives a http upgrade request from a client, completes this request
// and establishes the websocket connection. It then opens up concurent read/write
// listener for the user.
func (server *WsServer) ServeWs(w http.ResponseWriter, r *http.Request) {
	wsConnection, err := connection.UpgradeHTTPToWS(w, r)
	if err != nil {
		fmt.Println("Error in establishing websocket connection: ", err)
	}

	// Create user -- TODO: change from testUser to something else
	user := CreateUser("testUser", wsConnection, server)
	log.Printf("[INFO] new client connected")

	go user.CircularWrite(setting.WsServerSetting.Ping, setting.WsServerSetting.MaxWriteWaitTime)
	go user.CircularRead(setting.WsServerSetting.MaxMessageSize, setting.WsServerSetting.Pong)

	server.register <- user
}
