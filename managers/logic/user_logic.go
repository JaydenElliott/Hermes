package logic

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type User struct {
	userId     *string
	username   *string
	channels   map[*Channel]bool
	threads    map[*Thread]bool
	conn       *websocket.Conn
	wsServer   *WsServer
	dataBuffer chan []byte
}

// Create user method -> Used by user_manager.go
func CreateUser(userName string, conn *websocket.Conn, wsServer *WsServer) *User {
	userID := uuid.New().String()
	channels := make(map[*Channel]bool)
	threads := make(map[*Thread]bool)
	return &User{&userID, &userName, channels, threads, conn, wsServer, make(chan []byte, 256)}
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

func (user *User) GetConn() *websocket.Conn {
	return user.conn
}

func (user *User) GetWsSever() *WsServer {
	return user.wsServer
}

func (user *User) DisconnectWithWsServer() error {
	user.wsServer.unregister <- user
	close(user.dataBuffer)
	err := user.conn.Close()
	if err != nil {
		return fmt.Errorf("[Error] unable to disconnect user %s with the ws server", *user.GetUsername())
	}
	return nil
}
