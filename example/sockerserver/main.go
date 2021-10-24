package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/coolstina/fishsocket"
)

func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		cancel()
	}()

	server := fishsocket.NewServer(":8090")
	server.ServeMux().HandleFunc("/websocket", server.DefaultEchoTest())
	if err := server.Run(ctx); err != nil {
		log.Printf("failed to websocker serve:%v\n", err)
	}
}
