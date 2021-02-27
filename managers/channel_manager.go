package managers

import (
	"arcstack/arcstack-chat-server/managers/logic/channel"
)

// Create new Channel using CreateChanel_ parameters
func CreateChannel(p CreateChannel_) *channel.Channel {
	return channel.Create(p.ChannelName)
}
