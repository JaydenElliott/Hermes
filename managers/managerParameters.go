package managers

import (
	"arcstack/arcstack-chat-server/managers/logic"
	"github.com/gorilla/websocket"
)

// Parameters for CreateChannel function
type CreateChannel_ struct {
	ChannelName string
}

type CreateUser_ struct {
	UserName string
	conn     *websocket.Conn
	wsServer *logic.WsServer
}
