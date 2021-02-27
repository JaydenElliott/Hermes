package logic

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type Channel struct {
	ChannelID   *string
	ChannelName *string
	Users       map[*User]bool
	Threads     map[*Thread]bool
}

// Get Channel ID
func (channel *Channel) GetID() *string {
	return channel.ChannelID
}

// Get Channel name
func (channel *Channel) GetName() *string {
	return channel.ChannelName
}

// Description:    Gets all users in a specific Channel.
// Input:          getUsersParams struct (logicParameters.go).
// Returns:        List of pointers to user username or userID and error.
func (channel *Channel) GetUsers(p GetUsersParams_) ([]*string, error) {
	var users []*string
	var errorMsg error = nil

	// Loop through and append to return array all users satisfying users: True
	for user, value := range channel.Users {
		if value {
			if p.ReturnType == "username" {
				users = append(users, user.username)
			} else if p.ReturnType == "userId" {
				users = append(users, user.userId)
			} else {
				errorMsg = errors.New(fmt.Sprintf("Invalid getChannelUsers() input parameter: %s", p.ReturnType))
				users = nil
				break
			}
		}
	}
	return users, errorMsg
}

// Create a new thread within a channel. Function must be called from an instantiated channel.
func (channel *Channel) CreateThread() *Thread {
	threadID := uuid.New().String()
	users := make(map[*User]bool)
	return &Thread{&threadID, users, channel}
}
