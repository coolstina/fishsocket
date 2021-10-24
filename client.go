package fishsocket

import (
	"log"
	"net/url"

	"github.com/gorilla/websocket"
)

type Client struct {
	conn *websocket.Conn
}

func (client *Client) Connect() *websocket.Conn {
	return client.conn
}

func NewClient(address string, path string, ops ...ClientOption) (*Client, error) {
	options := clientOptions{
		connectScheme: ConnectSchemeForWS,
	}

	for _, o := range ops {
		o.apply(&options)
	}

	// WebSocket url.
	schema := options.connectScheme.String()
	u := url.URL{Scheme: schema, Host: address, Path: path}

	// Dial WebSocket server.
	dial, response, err := websocket.DefaultDialer.Dial(
		u.String(),
		options.requestHeader,
	)
	if err != nil {
		return nil, err
	}

	log.Printf("dail websocket server response: %+v\n", response)
	return &Client{conn: dial}, nil
}
