package logic

import "github.com/google/uuid"

type User struct {
	userId   *string
	username *string
	channels map[*Channel]bool
	threads  map[*Thread]bool
}

// Create user method -> Used by user_manager.go
func CreateUser(userName string) *User {
	userID := uuid.New().String()
	channels := make(map[*Channel]bool)
	threads := make(map[*Thread]bool)
	return &User{&userID, &userName, channels, threads}
}

func (user *User) GetID() *string {
	return user.userId
}

func (user *User) GetUsername() *string {
	return user.username
}

func (user *User) GetChannels() map[*Channel]bool {
	return user.channels
}

func (user *User) GetThreads() map[*Thread]bool {
	return user.threads
}
