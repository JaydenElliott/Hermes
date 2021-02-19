package managers

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
	users         []*User
}

type User struct {
	userId   *string
	email    *string
	username *string
	channels []*Channel // chat channels this user is apart of  ** REVIEW (would a map be better)
}

// Handles all business logic relating to a Channel
type ChannelManager struct {
	channelManagerID *string
	channels         []*Channel
}

type Channel struct {
	channelID   *string
	channelName *string
	users       []*User
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
