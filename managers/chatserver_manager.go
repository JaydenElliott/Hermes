package managers

import (
	"arcstack/arcstack-chat-server/managers/logic"
	"flag"
	"fmt"
	"net/http"
)

/*
Define Fundamental Types
*/

// Struct that bundles and acts as a controller between the users and channels
type ChatServerManager struct {
	UserManager    *UserManager
	ChannelManager *ChannelManager
	wsServer       *logic.WsServer
}

// Handles all business logic relating to a User
type UserManager struct {
	userManagerId *string
	users         []*logic.User
}

// Handles all business logic relating to a Channel
type ChannelManager struct {
	channelManagerID *string
	channels         []*logic.Channel
}

// Initialises managers and defines object design pattern`
func InitialiseManager() *ChatServerManager {

	controller := new(ChatServerManager)

	// Initialise the websocketServer
	server := logic.NewWsServer()
	controller.wsServer = server

	// Initialise child structs
	um := new(UserManager)
	cm := new(ChannelManager)

	controller.UserManager = um
	controller.ChannelManager = cm

	return controller

}

func (chatManager *ChatServerManager) RunWsServer() {
	// TODO: move this to config
	var addr = flag.String("addr", ":8080", "http service address")
	// Start websocket register listener
	go chatManager.wsServer.Run()
	// Start websocket read/write pump listening
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		chatManager.wsServer.ServeWs(w, r)
	})
	// Port listening
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		fmt.Println("ListenAndServe Error: ", err)
	}
}
