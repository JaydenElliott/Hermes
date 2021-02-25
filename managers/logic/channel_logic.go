package logic

type Channel struct {
	channelID   string
	channelName string
	users       map[*User]bool
	threads     map[*Thread]bool
}

// Get Channel ID
func (channel *Channel) getID() string {
	return channel.channelID
}

// Get Channel name
func (channel *Channel) getName() string {
	return channel.channelName
}

// Parameters for getUsers function
type getUsersParams struct {
	returnType string // [id, username]
}

// Gets all users in a channel
func (channel *Channel) getUsers(p getUsersParams) []string {
	var users []string
	for user, value := range channel.users {
		if value {
			if p.returnType == "username" {
				users = append(users, user.username)
			} else if p.returnType == "id" {
				users = append(users, user.userId)
			}

		}
	}
	return users
}
