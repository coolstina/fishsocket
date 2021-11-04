package fishsocket

import (
	"context"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/coolstina/fishserver"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestClientSuite(t *testing.T) {
	suite.Run(t, new(ClientSuite))
}

type ClientSuite struct {
	suite.Suite
	client *Client
}

const socketPath = "/web"
const socketHost = "localhost:9090"

func (suite *ClientSuite) BeforeTest(suiteName, testName string) {
	go func() {
		serve()
	}()

	suite.client = NewClient(socketHost,
		WithClientReconnectInterval(1*time.Second),
	).SetPath(socketPath)
	assert.NotNil(suite.T(), suite.client)
}

func (suite *ClientSuite) Test_Connect() {
	connect := suite.client.Connect()
	err := connect.WriteMessage(1, []byte("helloworld"))
	assert.NoError(suite.T(), err)

	mt, message, err := connect.ReadMessage()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 1, mt)
	assert.Equal(suite.T(), []byte("helloworld"), message)
}

func serve() {
	mux := http.NewServeMux()
	mux.HandleFunc("/websocket", echo)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	server := fishserver.NewServer(socketHost,
		fishserver.WithWaitTimeout(5*time.Millisecond),
		fishserver.WithContext(ctx),
		fishserver.WithCancelFunc(cancel),
	).SetHandler(mux)

	if err := server.Run(); err != nil {
		log.Printf("failed to serve: %+v\n", err)
	}
}

func echo(writer http.ResponseWriter, request *http.Request) {
	upgrader := UpgraderHandleFuncWithDefault()
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
	}
}
