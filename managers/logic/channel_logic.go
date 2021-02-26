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
}

// Create new Channel using CreateChanel_ parameters
func (channel *Channel) CreateChannel(p CreateChannel_) *Channel {
	channelID := uuid.New().String() // generate channel uuid
	users := make(map[*User]bool)
	threads := make(map[*Thread]bool)
	return &Channel{&channelID, &p.ChannelName, users, threads}
}

// Get Channel ID
func (channel *Channel) getID() *string {
	return channel.channelID
}

// Get Channel name
func (channel *Channel) getName() *string {
	return channel.channelName
}

// Description:    Gets all users in a specific Channel.
// Input:          getUsersParams struct (logicParameters.go).
// Returns:        List of pointers to user username or userID and error.
func (channel *Channel) getChannelUsers(p GetUsersParams_) ([]*string, error) {
	var users []*string
	var errorMsg error = nil

	// Loop through and append to return array all users satisfying users: True
	for user, value := range channel.users {
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
