package models

// Obtain class attributes
type Channel interface {
	getID() *string
	getName() *string
	getUsers() ([]*string, error)
}

// Interact with the channel class
type ChannelLogic interface {
	CreateChannel(channel Channel)
	RemoveChannel(channel Channel)
	MergeChannel(channel Channel)
	FindChannel(name string)
}
