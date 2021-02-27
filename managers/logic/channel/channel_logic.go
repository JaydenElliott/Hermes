package channel

import (
	"arcstack/arcstack-chat-server/managers/logic/thread"
	"arcstack/arcstack-chat-server/managers/logic/user"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type Channel struct {
	channelID   *string
	channelName *string
	users       map[*user.User]bool
	threads     map[*thread.Thread]bool
}

// Create channel method -> Used by channel_manager.go
// Created here to allow Channel to be immutable
func Create(channelName string) *Channel {
	channelID := uuid.New().String() // generate channel uuid
	users := make(map[*user.User]bool)
	threads := make(map[*thread.Thread]bool)
	return &Channel{&channelID, &channelName, users, threads}
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

func (channel *Channel) GetAllUsers() map[*user.User]bool {
	return channel.users

}

func (channel *Channel) getThreads() map[*thread.Thread]bool {
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
//func (channel *Channel) CreateThread() *thread.Thread {
//	threadID := uuid.New().String()
//	users := make(map[*user.User]bool)
//	return &thread.Thread{&threadID, users, channel}
//}
