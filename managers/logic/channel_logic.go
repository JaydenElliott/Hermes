package logic

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type Channel struct {
	channelID   *string
	channelName *string
	users       map[*User]bool
	threads     map[*Thread]bool
	register    chan *User
	unregister  chan *User
	broadcast   chan *Message
}

// Create channel method -> Used by channel_manager.go
// Created here to allow Channel to be immutable
func CreateChannel(channelName string) *Channel {

	// Initialise fields
	channelID := uuid.New().String()
	users := make(map[*User]bool)
	threads := make(map[*Thread]bool)
	register := make(chan *User)
	unregister := make(chan *User)
	broadcast := make(chan *Message)

	return &Channel{&channelID,
		&channelName,
		users,
		threads,
		register,
		unregister,
		broadcast}
}

// Adds a user to a room
func (channel *Channel) registerUser(user *User) {

	// Send join message to room

	// Register user
	channel.users[user] = true
}

// Removes user from a room
func (channel *Channel) unregisterUser(user *User) {
	// Send leave message to room

	// Remove from room
	if _, ok := channel.users[user]; ok {
		delete(channel.users, user)
	}

}

func (channel *Channel) broadcastToUsers() {
}

/*
	Methods to get channel fields
*/

func (channel *Channel) GetID() *string {
	return channel.channelID
}

func (channel *Channel) GetName() *string {
	return channel.channelName
}

func (channel *Channel) GetAllUsers() map[*User]bool {
	return channel.users

}

func (channel *Channel) getThreads() map[*Thread]bool {
	return channel.threads
}

// Description:    Gets all users in a specific Channel.
// Input:          getUsersParams struct (logicParameters.go).
// Returns:        List of pointers to user username or userID and error.
func (channel *Channel) GetUsers(p GetUsersParams_) ([]*string, error) {
	var users []*string
	var errorMsg error = nil

	// Loop through and append to return array all users satisfying users: True
	for User, value := range channel.users {
		if value {
			if p.ReturnType == "username" {
				users = append(users, User.GetUsername())
			} else if p.ReturnType == "userId" {
				users = append(users, User.GetID())
			} else {
				errorMsg = errors.New(fmt.Sprintf("Invalid getChannelUsers() input parameter: %s", p.ReturnType))
				users = nil
				break
			}
		}
	}
	return users, errorMsg
}

/*
	Channel modification methods
*/

func (channel *Channel) UpdateName(p UpdateName_) {
	channel.channelName = &p.UpdatedName
}

/*
	Channel Threads Methods
*/

// Create a new thread within a channel. Function must be called from an instantiated channel.
func (channel *Channel) CreateThread() *Thread {
	threadID := uuid.New().String()
	users := make(map[*User]bool)
	return &Thread{&threadID, users, channel}
}
