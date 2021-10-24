package fishsocket

import "time"

// ServerOption WebSocket server options interface.
type ServerOption interface {
	apply(*serverOptions)
}

type serverOptions struct {
	waitTimeout time.Duration
}

func WithServerShutdownWaitTimeout() {

}
