package fishsocket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

// UpgraderHandleFunc WebSocket upgrader handler func.
type UpgraderHandleFunc func(w http.ResponseWriter,
	r *http.Request, responseHeader http.Header) (*websocket.Conn, error)

// UpgraderHandleFuncWithDefault Get a new default upgrader handler func type.
func UpgraderHandleFuncWithDefault() UpgraderHandleFunc {
	return (&websocket.Upgrader{}).Upgrade
}
