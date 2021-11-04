package fishsocket

import (
	"fmt"
	"log"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

type Client struct {
	options *clientOptions
	lock    sync.Mutex
	url     *url.URL
	conn    *websocket.Conn
	err     error // Automatic reconnection in case of error.
}

func (client *Client) Connect() *websocket.Conn {
	client.lock.Lock()
	defer client.lock.Unlock()

	if client.conn == nil || client.err != nil {
		ul := client.url.String()
		log.Printf("connect to websocket [%s] ...\n", ul)

		// Dial WebSocket server.
		dial, resp, err := websocket.DefaultDialer.Dial(ul,
			client.options.requestHeader,
		)

		// Reconnect to WebSocket Server.
		if err != nil {
			for {
				log.Printf("reconnect to websocket [%s] ...\n", ul)
				dial, resp, err = websocket.DefaultDialer.Dial(ul,
					client.options.requestHeader,
				)

				if err != nil {
					time.Sleep(client.options.interval)
					continue
				}

				log.Printf("reconnect to websocket [%s] successfully~\n", ul)
				break
			}
		}

		log.Printf("websocket connected response [%+v]\n", resp)
		client.conn = dial
		client.err = nil
	}

	return client.conn
}

func (client *Client) SetConnectError(err error) {
	client.lock.Lock()
	defer client.lock.Unlock()

	client.err = err
}

func (client *Client) SetPath(path string) *Client {
	client.lock.Lock()
	defer client.lock.Unlock()

	if path != "" {
		if !strings.HasPrefix(path, "/") {
			path = fmt.Sprintf("/%s", path)
		}
		client.url.Path = path
	}

	return client
}

func NewClient(address string, ops ...ClientOption) *Client {
	options := &clientOptions{
		connectScheme: ConnectSchemeForWS,
		interval:      10 * time.Second,
	}

	for _, o := range ops {
		o.apply(options)
	}

	// WebSocket url.
	u := &url.URL{
		Scheme: options.connectScheme.String(),
		Host:   address,
		Path:   "/",
	}

	return &Client{options: options, url: u}
}
