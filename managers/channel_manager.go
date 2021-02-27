package managers

import (
	"arcstack/arcstack-chat-server/managers/logic"
)

// Create new Channel using CreateChanel_ parameters
func CreateChannel(p CreateChannel_) *logic.Channel {
	return logic.CreateChannel(p.ChannelName)
}
