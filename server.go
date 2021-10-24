package socket

import (
	"context"
	"log"
	"net/http"
	"time"
)

type Server struct {
	mux     *http.ServeMux
	address string
}

func (s *Server) ServeMux() *http.ServeMux {
	return s.mux
}

func NewServer(address string) *Server {
	srv := &Server{
		mux:     http.NewServeMux(),
		address: address,
	}

	return srv
}

func (s *Server) DefaultEchoTest() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
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
}

func (s *Server) Run(ctx context.Context, ops ...ServerOption) (err error) {
	options := serverOptions{
		waitTimeout: 5 * time.Second,
	}

	for _, o := range ops {
		o.apply(&options)
	}

	srv := &http.Server{
		Addr:    s.address,
		Handler: s.mux,
	}

	// Server run.
	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("server started at %s\n", s.address)
	<-ctx.Done()

	// Wait
	timeout, cancel := context.WithTimeout(context.Background(), options.waitTimeout)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(timeout); err != nil {
		log.Printf("server shutdown failed: %s\n", err)
	}
	log.Printf("server exited properly\n")

	if err == http.ErrServerClosed {
		return nil
	}

	return err
}
