package socket

import (
	"net/http"

	"github.com/gorilla/websocket"
)

type UpgraderHandleFunc func(w http.ResponseWriter, r *http.Request, responseHeader http.Header) (*websocket.Conn, error)

func UpgraderHandleFuncWithDefault() UpgraderHandleFunc {
	upgrader := websocket.Upgrader{}
	return upgrader.Upgrade
}
