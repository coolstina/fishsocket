package main

import (
	"log"
	"net/http"
	"time"

	"github.com/coolstina/fishserver"
	"github.com/coolstina/fishsocket"
)

const address = "localhost:9099"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/websocket", echo)

	server := fishserver.NewServer(address,
		fishserver.WithWaitTimeout(5*time.Millisecond),
	).SetHandler(mux)

	if err := server.Run(); err != nil {
		log.Printf("failed to serve: %+v\n", err)
	}
}

func echo(writer http.ResponseWriter, request *http.Request) {
	upgrader := fishsocket.UpgraderHandleFuncWithDefault()
	conn, err := upgrader(writer, request, nil)
	if err != nil {
		log.Printf("websocket server upgrader failed: %+v\n", err)
		return
	}

	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("websocket server read message to failed: %+v\n", err)
			break
		}

		log.Printf("websocker server recevied message: %s\n", message)
		err = conn.WriteMessage(mt, message)
		if err != nil {
			log.Printf("websocket server write message to failed: %+v\n", err)
		}

		log.Printf("websocket server send message: %s\n", message)
		time.Sleep(5 * time.Second)
	}
}
