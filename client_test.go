package socket

import (
	"context"
	"log"
	"os"
	"os/signal"
	"testing"
	"time"

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

const socketPath = "/websocket"
const socketHost = "localhost:8970"

func (suite *ClientSuite) BeforeTest(suiteName, testName string) {
	go func() {
		socketServe()
	}()

	var err error
	suite.client, err = NewClient(
		socketHost,
		socketPath,
	)
	assert.NoError(suite.T(), err)
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

func socketServe() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		cancel()
	}()

	server := NewServer(socketHost)
	server.ServeMux().HandleFunc(socketPath, server.DefaultEchoTest())

	if err := server.Run(ctx); err != nil {
		log.Printf("failed to serve:+%v\n", err)
	}
}
