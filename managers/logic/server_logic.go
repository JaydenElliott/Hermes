package logic

import (
	"arcstack/arcstack-chat-server/pkg/util/connection"
	"fmt"
	"net/http"
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
		}
	}
}

func (server *WsServer) ServeHome(w http.ResponseWriter, r *http.Request) {
	// Check for the correct endpoint location.
	if r.URL.Path != "/" {
		http.Error(w, "Endpoint not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Request not supported on endpoint `/` -> Only GET request allowed", http.StatusMethodNotAllowed)
	}
	http.ServeFile(w, r, "home.html") // change this to the users html endpoint // TODO
}

func (server *WsServer) ServeWs(w http.ResponseWriter, r *http.Request) {
	wsConnection, err := connection.UpgradeHTTPToWS(w, r)
	if err != nil {
		fmt.Println("Error in establishing websocket connection: ", err)
	}

	// TODO setup user (client) side connection with read/write pump

}
