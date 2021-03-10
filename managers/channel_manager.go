package managers

import (
	"arcstack/arcstack-chat-server/managers/logic"
)

// Create new Channel using CreateChanel_ parameters
func (cm *ChannelManager) CreateChannel(p CreateChannel_) *logic.Channel {
	channel := logic.CreateChannel(p.ChannelName)
	go channel.Run()
	cm.channels[channel] = true
	return channel
}
