package main

import (
	"fmt"
	"log"
	"time"

	"github.com/coolstina/fishsocket"
	"github.com/gorilla/websocket"
)

const (
	address = "localhost:9099"
	path    = "/websocket"
)

func main() {
	client := fishsocket.NewClient(address,
		fishsocket.WithClientReconnectInterval(1*time.Second),
	).SetPath(path)

	for {
		// Write message.
		err := client.Connect().WriteMessage(websocket.TextMessage, []byte("Hello WebSocket Server"))
		if err != nil {
			log.Printf("Write message to WebSocket Server to failed")
			client.SetConnectError(err)
			continue
		}

		// Read message.
		mt, message, err := client.Connect().ReadMessage()
		if err != nil {
			fmt.Printf("Read message from WebSocket client to failed: %+v\n", err)
			client.SetConnectError(err)
			continue
		}

		fmt.Printf("mt: %+v, message: %s\n", mt, message)
	}
}
