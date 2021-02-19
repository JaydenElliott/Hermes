package models

// Obtain class attributes
type Channel interface {
	getID() string
	getName() string
	getUsers() []string
}

// Interact with the channel class
type ChannelLogic interface {
	AddChannel(channel Channel)
	RemoveChannel(channel Channel)
	MergeChannel(channel Channel)
	FindChannel(name string)
}
