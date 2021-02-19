package managers

/* Define Fundamental Types */

// Struct that bundles and acts as a controller between the users and channels
type ChatServerManager struct {
	UserManager    *UserManager
	ChannelManager *ChannelManager
}

// Handles all business logic relating to a User
type UserManager struct {
	users []*User
}

type User struct {
	email       *string
	displayName *string
	channels    []*Channel // chat channels this user is apart of  ** REVIEW (would a map be better)
}

// Handles all business logic relating to a Channel
type ChannelManager struct {
}

type Channel struct {
}
