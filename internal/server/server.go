package server

import (
	"context"
	"net/http"
)

type server struct {
	httpServer http.Server
}

func newServer(bindAddress string, handler http.Handler) *server {
	server := &server{
		httpServer: http.Server{
			Addr:    bindAddress,
			Handler: handler,
		},
	}
	return server
}

func (server *server) ListenAndServe() {
	server.httpServer.ListenAndServe()
}

func (server *server) Shutdown(context context.Context) error {
	return server.httpServer.Shutdown(context)
}
