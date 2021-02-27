package managers

import (
	"arcstack/arcstack-chat-server/managers/logic/channel"
	"arcstack/arcstack-chat-server/managers/logic/user"
)

/*
Define Fundamental Types
*/

// Struct that bundles and acts as a controller between the users and channels
type ChatServerManager struct {
	UserManager    *UserManager
	ChannelManager *ChannelManager
}

// Handles all business logic relating to a User
type UserManager struct {
	userManagerId *string
	users         []*user.User
}

// Handles all business logic relating to a Channel
type ChannelManager struct {
	channelManagerID *string
	channels         []*channel.Channel
}

// Initialises managers and defines object design pattern`
func InitialiseManager() *ChatServerManager {
	controller := new(ChatServerManager)

	// TODO: run config files here to ensure all databases and servers are running

	// Initialise child structs
	um := new(UserManager)
	cm := new(ChannelManager)

	controller.UserManager = um
	controller.ChannelManager = cm

	return controller

}
