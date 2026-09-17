package main

import (
	"context"
	"errors"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(
	port string,
	handler http.Handler,
) *Server {

	httpServer := &http.Server{
		Addr:    ":" + port,
		Handler: handler,

		ReadHeaderTimeout: 5 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      10 * time.Second,
		IdleTimeout:       60 * time.Second,
	}

	return &Server{
		httpServer: httpServer,
	}
}

func (s *Server) Run() error {

	err := s.httpServer.ListenAndServe()

	if err != nil && !errors.Is(
		err,
		http.ErrServerClosed,
	) {
		return err
	}

	return nil
}

func (s *Server) Shutdown(
	ctx context.Context,
) error {

	return s.httpServer.Shutdown(ctx)
}
