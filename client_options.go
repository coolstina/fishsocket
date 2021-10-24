package socket

import "net/http"

// ClientOption WebSocket client options interface.
type ClientOption interface {
	apply(*clientOptions)
}

type clientOptions struct {
	requestHeader http.Header
	connectScheme ConnectSchema
}

func WithClientRequestHeader(header http.Header) ClientOption {
	return clientOptionsFunc(func(options *clientOptions) {
		options.requestHeader = header
	})
}

func WithClientConnectSchema(schema ConnectSchema) ClientOption {
	return clientOptionsFunc(func(options *clientOptions) {
		options.connectScheme = schema
	})
}
