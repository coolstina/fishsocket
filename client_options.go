package fishsocket

import (
	"net/http"
	"time"
)

// ClientOption WebSocket client options interface.
type ClientOption interface {
	apply(*clientOptions)
}

type clientOptions struct {
	path          time.Duration
	interval      time.Duration
	requestHeader http.Header
	connectScheme ConnectSchema
}

// WithClientRequestHeader Specific client request header.
func WithClientRequestHeader(header http.Header) ClientOption {
	return clientOptionsFunc(func(options *clientOptions) {
		options.requestHeader = header
	})
}

// WithClientConnectSchema Specific client connect schema,
// Optional values include ConnectSchemeForWS(ws) or ConnectSchemeForWSS(wss).
func WithClientConnectSchema(schema ConnectSchema) ClientOption {
	return clientOptionsFunc(func(options *clientOptions) {
		options.connectScheme = schema
	})
}

// WithClientReconnectInterval Specific WebSocket Client
// reconnection interval to WebSocket Server.
func WithClientReconnectInterval(interval time.Duration) ClientOption {
	return clientOptionsFunc(func(options *clientOptions) {
		options.interval = interval
	})
}
