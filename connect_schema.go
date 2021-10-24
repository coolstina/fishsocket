package fishsocket

type ConnectSchema string

func (c ConnectSchema) String() string {
	return string(c)
}

const (
	ConnectSchemeForWS  ConnectSchema = "ws"
	ConnectSchemeForWSS ConnectSchema = "wss"
)
