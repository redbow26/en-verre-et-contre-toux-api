package server

import (
	"fmt"
	"os"
	"os/signal"
)

func (server *Server) Listen(addr string) error {
	return server.app.Listen(addr)
}

func (server *Server) Shutdown() error {
	return server.app.Shutdown()
}

func (server *Server) Graceful() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = server.Shutdown()
	}()
}

func (server *Server) ListenWithGraceful(addr string) error {
	server.Graceful()
	return server.Listen(addr)
}
