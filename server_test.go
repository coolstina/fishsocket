package socket

import (
	"context"
	"log"
	"os"
	"os/signal"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

func TestServerSuite(t *testing.T) {
	suite.Run(t, new(ServerSuite))
}

type ServerSuite struct {
	suite.Suite
	client *Client
}

func (suite *ServerSuite) BeforeTest(suiteName, testName string) {
	suite.client = nil
}

func (suite *ServerSuite) Test_NewServer() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		cancel()
	}()

	server := NewServer("localhost:8969")
	if err := server.Run(ctx); err != nil {
		log.Printf("failed to serve:+%v\n", err)
	}
}
