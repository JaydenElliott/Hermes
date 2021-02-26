package logic

import (
	"errors"
	"fmt"
)

type Thread struct {
	threadID *string
	users    map[*User]bool
	channel  *Channel
}

// Get Thread ID
func (thread *Thread) getID() *string {
	return thread.threadID
}

// Get parent channel
func (thread *Thread) getParentChannel() *Channel {
	return thread.channel
}

// Description:    Gets all users in a specific Thread.
// Input:          getUsersParams struct (logicParameters.go).
// Returns:        List of pointers to user username or userID and error.
func (thread *Thread) getThreadUsers(p GetUsersParams_) ([]*string, error) {
	var users []*string
	var errorMsg error = nil

	// Loop through and append to return array all users satisfying users: True
	for user, value := range thread.users {
		if value {
			if p.ReturnType == "username" {
				users = append(users, user.username)
			} else if p.ReturnType == "userId" {
				users = append(users, user.userId)
			} else {
				errorMsg = errors.New(fmt.Sprintf("Invalid getThreadUsers() input parameter: %s", p.ReturnType))
				users = nil
				break
			}
		}

	}
	return users, errorMsg
}
