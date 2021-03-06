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

func (user *User) configureConn(maxMessageSize int64, pong time.Duration) {
	user.conn.SetReadLimit(maxMessageSize)
	_ = user.conn.SetReadDeadline(time.Now().Add(pong))
	user.conn.SetPongHandler(func(string) error { _ = user.conn.SetReadDeadline(time.Now().Add(pong)); return nil })
}

func (user *User) CircularRead(maxMessageSize int64, pong time.Duration) {
	defer func() {
		_ = user.DisconnectWithWsServer()
	}()

	user.configureConn(maxMessageSize, pong)

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

func (user *User) CircularWrite(period time.Duration, maxWriteWaitTime time.Duration) {
	ticker := time.NewTicker(period)
	defer func() {
		ticker.Stop()
		_ = user.conn.Close()
	}()

	for {
		select {
		case message, ok := <-user.dataBuffer:
			// SetWriteDeadline sets the maxWriteWaitTime as a deadline on the underlying network connection.
			// If maxWriteWaitTime has timed out, the websocket state is corrupt and
			// all future writes will return an error :(
			//A zero value for t means writes will not time out.
			_ = user.conn.SetWriteDeadline(time.Now().Add(maxWriteWaitTime))
			if !ok {
				// The WsServer closed the channel.
				_ = user.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := user.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			_, err = w.Write(message)
			if err != nil {
				return
			}

			// Attach queued chat messages to the current websocket message.
			n := len(user.dataBuffer)
			for i := 0; i < n; i++ {
				_, _ = w.Write(newline)
				_, _ = w.Write(<-user.dataBuffer)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			_ = user.conn.SetWriteDeadline(time.Now().Add(maxWriteWaitTime))
			if err := user.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
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
