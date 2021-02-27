package thread

import (
	"arcstack/arcstack-chat-server/managers/logic/channel"
	"arcstack/arcstack-chat-server/managers/logic/user"
	"errors"
	"fmt"
)

type Thread struct {
	threadID *string
	users    map[*user.User]bool
	channel  *channel.Channel
}

// Get Thread ID
func (thread *Thread) GetID() *string {
	return thread.threadID
}

// Get parent channel
func (thread *Thread) GetParentChannel() *channel.Channel {
	return thread.channel
}

// Description:    Gets all users in a specific Thread.
// Input:          getUsersParams struct (logicParameters.go).
// Returns:        List of pointers to user username or userID and error.
func (thread *Thread) GetThreadUsers(p GetUsersParams_) ([]*string, error) {
	var users []*string
	var errorMsg error = nil

	// Loop through and append to return array all users satisfying users: True
	for User, value := range thread.users {
		if value {
			if p.ReturnType == "username" {
				users = append(users, User.GetUsername())
			} else if p.ReturnType == "userId" {
				users = append(users, User.GetID())
			} else {
				errorMsg = errors.New(fmt.Sprintf("Invalid getThreadUsers() input parameter: %s", p.ReturnType))
				users = nil
				break
			}
		}

	}
	return users, errorMsg
}
