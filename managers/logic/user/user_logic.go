package user

import (
	"arcstack/arcstack-chat-server/managers/logic/channel"
	"arcstack/arcstack-chat-server/managers/logic/thread"
)

type User struct {
	userId   *string
	username *string
	channels map[*channel.Channel]bool
	threads  map[*thread.Thread]bool
}

func (user *User) getID() *string {
	return user.userId
}

func (user *User) getUsername() *string {
	return user.username
}

func (user *User) getChannels() map[*channel.Channel]bool {
	return user.channels
}

func (user *User) getThreads() map[*thread.Thread]bool {
	return user.threads
}
