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

func (user *User) GetID() *string {
	return user.userId
}

func (user *User) GetUsername() *string {
	return user.username
}

func (user *User) GetChannels() map[*channel.Channel]bool {
	return user.channels
}

func (user *User) GetThreads() map[*thread.Thread]bool {
	return user.threads
}
