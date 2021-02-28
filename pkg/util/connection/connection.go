package connection

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// BufferSizes default:
// readBufferSize: 4096
// writeBufferSize: 4096
type BufferSizes struct {
	readBufferSize  int
	writeBufferSize int
}

//makeUpgrader: make a websocket upgrader based on specified buffersizes.
func makeUpgrader(bufferSizes *BufferSizes) websocket.Upgrader {

	if bufferSizes.readBufferSize == 0 {
		bufferSizes.readBufferSize = 4096
	}
	if bufferSizes.writeBufferSize == 0 {
		bufferSizes.writeBufferSize = 4096
	}

	var upgrader = websocket.Upgrader{
		ReadBufferSize:  bufferSizes.readBufferSize,
		WriteBufferSize: bufferSizes.writeBufferSize,
	}
	return upgrader
}

// upgradeHTTPToWS upgrades the HTTP server connection to the WebSocket protocol.
func upgradeHTTPToWS(upgrader websocket.Upgrader, w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return nil, err
	}
	return conn, err
}
