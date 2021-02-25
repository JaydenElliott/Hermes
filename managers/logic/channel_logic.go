package logic

type Channel struct {
	channelID   string
	channelName string
	users       map[*User]bool
	threads     map[*Thread]bool
}

func (channel *Channel) getID() string {
	return channel.channelID
}

func (channel *Channel) getName() string {
	return channel.channelName
}

type getUsersParams struct {
	returnType string
}

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
