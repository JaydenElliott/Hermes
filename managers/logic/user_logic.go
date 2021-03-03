package logic

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"time"
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

// Conn default:
// maxMessageSize: 1000
// duration: 60 * time.Second
type Conn struct {
	// Maximum message size allowed from peer.
	maxMessageSize int64
	// Max time till next pong from peer
	duration time.Duration
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

func (user *User) configureConn(maxMessageSize int64, duration time.Duration) {

	if maxMessageSize == 0 {
		maxMessageSize = 10000
	}

	if duration == 0 {
		duration = 60 * time.Second
	}

	user.conn.SetReadLimit(maxMessageSize)
	_ = user.conn.SetReadDeadline(time.Now().Add(duration))
	user.conn.SetPongHandler(func(string) error { _ = user.conn.SetReadDeadline(time.Now().Add(duration)); return nil })
}

func (user *User) CircularRead(maxMessageSize int64, duration time.Duration) {
	defer func() {
		_ = user.DisconnectWithWsServer()
	}()

	user.configureConn(maxMessageSize, duration)

	// Start endless read loop, waiting for messages from client
	for {
		_, jsonMessage, err := user.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("[ERROR] unexpected close error: %v", err)
			}
			break
		}
		user.wsServer.broadcast <- jsonMessage
	}
}

func (user *User) DisconnectWithWsServer() error {
	user.wsServer.unregister <- user
	close(user.dataBuffer)
	err := user.conn.Close()
	if err != nil {
		return fmt.Errorf("[ERROR] unable to disconnect user %s with the ws server", *user.GetUsername())
	}
	return nil
}
