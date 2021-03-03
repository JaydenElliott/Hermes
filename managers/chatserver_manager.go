package managers

import (
	"arcstack/arcstack-chat-server/managers/logic"
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
	go chatManager.wsServer.Run()
}
