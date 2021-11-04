![fishsocket](assert/banner/fishsocket.jpg)

Use [gorilla/websocket](https://github.com/gorilla/websocket) quick build application.

## Installation

```shell script
go get -u github.com/coolstina/fishsocket
```


## Build WebSocket Server

```go
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
```

## Build WebSocket Client 

```go
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
```