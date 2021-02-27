package managers

import (
	"arcstack/arcstack-chat-server/managers/logic"
	"github.com/google/uuid"
)

// Create new Channel using CreateChanel_ parameters
func CreateChannel(p logic.CreateChannel_) *logic.Channel {
	channelID := uuid.New().String() // generate channel uuid
	users := make(map[*logic.User]bool)
	threads := make(map[*logic.Thread]bool)
	return &logic.Channel{&channelID, &p.ChannelName, users, threads}
}
