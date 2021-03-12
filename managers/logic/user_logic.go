package logic

import (
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"log"
	"time"
)

type User struct {
	UserId     string  `json:"UserId"` // encoded to be parsed with messages
	username   *string // name to be displayed around the server
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
	return &User{userID, &userName, channels, threads, conn, wsServer, make(chan []byte, 256)}
}

func (user *User) GetID() string {
	return user.UserId
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

// Configure read size limits / time out duration / pong handling
func (user *User) configureConn(maxMessageSize int64, pong time.Duration) {
	user.conn.SetReadLimit(maxMessageSize)
	_ = user.conn.SetReadDeadline(time.Now().Add(pong))
	user.conn.SetPongHandler(func(string) error { _ = user.conn.SetReadDeadline(time.Now().Add(pong)); return nil })
}

// CircularRead will continuously read incoming messages to the websocket.
//
// Parameters:
// 		maxMessageSize (int64) size limit of the message in bytes
// 		pong (time.Duration) set pong max response time to detect dead client
func (user *User) CircularRead(maxMessageSize int64, pong time.Duration) {

	// Disconnect websocket server
	defer func() {
		err := user.DisconnectWithWsServer()
		if err != nil {
			log.Printf("[ERROR] unexpected error when user trying to disconnect with the WebSocket error: %v", err)
		}
	}()

	// Set Read Timeout / Size Limits / Pong Handling
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

// CircularWrite handles sending messages to the connected user.
func (user *User) CircularWrite(ping time.Duration, maxWriteWaitTime time.Duration) {
	//  Define ticker to send client pings every "ping" duration.
	ticker := time.NewTicker(ping)

	// When the client side user connection is broken,
	// close the ws connection from the server side and log the error.
	defer func() {
		ticker.Stop()
		err := user.conn.Close()
		if err != nil {
			log.Printf("[ERROR] unexpected user connection close error: %v", err)
		}
	}()

	// Begin circular Write
	for {
		select {
		case message, ok := <-user.dataBuffer:
			// SetWriteDeadline sets the maxWriteWaitTime as a deadline on the underlying network connection.
			// If maxWriteWaitTime has timed out, the websocket state is corrupt and
			// all future writes will return an error :(
			// A zero value for t means writes will not time out.
			err := user.conn.SetWriteDeadline(time.Now().Add(maxWriteWaitTime))
			if err != nil {
				log.Printf("[ERROR] unexpected error for setting : %v", err)
			}
			if !ok {
				// The WsServer closed the channel.
				err := user.conn.WriteMessage(websocket.CloseMessage, []byte{})
				if err != nil {
					log.Printf("[ERROR] unexpected error when user connection writing messages: %v", err)
				}
				return
			}

			// Generate a writer for the next message to utilise
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
				_, err := w.Write(<-user.dataBuffer)
				if err != nil {
					log.Printf("[ERROR] unexpected error when writer writing data buffer: %v", err)
				}
			}
			if err := w.Close(); err != nil {
				return
			}

		// Every "ping" amount of time, ping client and wait for response.
		// No response => error.
		case <-ticker.C:
			// Set new write deadline
			err := user.conn.SetWriteDeadline(time.Now().Add(maxWriteWaitTime))
			if err != nil {
				log.Printf("[ERROR] unexpected error when user conection setting write deadline: %v", err)
			}
			// Send Ping
			if err := user.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				log.Printf("[ERROR] unexpected error when user conection writing messages: %v", err)
				return
			}
		}
	}
}

// DisconnectWithWsServer unregisters user from server
// closes the buffer channel and closes the websocket connection.
func (user *User) DisconnectWithWsServer() error {
	// Unregister user from websocket
	user.wsServer.unregister <- user

	// Close msg buffer channel
	close(user.dataBuffer)

	// Close websocket connection
	err := user.conn.Close()
	if err != nil {
		return err
	}
	return nil
}

func (user *User) newMessage(jsonMsg []byte) error {

	// Convert msg to the correct format
	msg := MessageUnmarshal(jsonMsg)
	if msg == nil {
		return errors.New("Unable to handle new message")
	}

	// User is the sender of the message
	msg.Sender = user

	switch msg.Action {
	case SendMessageAction:
		// Room to send message to
		channelName := msg.Target.GetName()

		// If channel exists, send the message to the channel's broadcast method
		if channel, _ := user.wsServer.FindChannel(FindChannelParams{channelName, nil}); channel != nil {
			channel.broadcast <- msg
		}

	case JoinChannelAction:
		//user.handleJoinChannelMessage(msg)

	case LeaveChannelAction:
		//user.handleLeaveChannelMessage(msg)
	}

	return nil
}

// HandleJoinChannel will add a user to a room or create a new
// room if no room with the associated name exists.
//
// Params message Message: the channel the user wants to join/create
func (user *User) HandleJoinChannel(message Message) {
	// Find the channel the user wants to join.
	// If channel doesn't exist, make one with associated ChannelName
	channel, _ := user.wsServer.FindChannel(FindChannelParams{&message.Message, nil})
	if channel == nil {
		channel = user.wsServer.NewWsChannel(message.Message)
	}

	// Append channel to users channel map
	user.channels[channel] = true

	// Add user to channel in the channel method
	channel.register <- user
}

// HandleLeaveChannel searches for the channel the user wants to leave
// and attempts to delete the user from the channel
func (user *User) HandleLeaveChannel(message Message) {

	// Find channel user requests to leave.
	channel, _ := user.wsServer.FindChannel(FindChannelParams{&message.Message, nil})

	// Attempt to delete channel from user map.
	ok := user.channels[channel]
	if ok {
		delete(user.channels, channel)
	} else {
		log.Println("Unable to remove user from channel as channel not found in user map.")
	}

	// Remove user from channel's users
	channel.unregister <- user
}
