package server

import (
	"context"
	"net/http"
	"time"

	"github.com/IskanderA1/handly/pkg/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(handler http.Handler, config config.Config) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           config.ServerAddress,
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			Handler:        handler,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
