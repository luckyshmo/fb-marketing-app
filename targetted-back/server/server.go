package server

import (
	"context"
	"net/http"
	"time"
)

// Backend is custom wrapper for additional functionality for our server
type Backend struct {
	httpServer *http.Server
}

func (s *Backend) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    100 * time.Second,
		WriteTimeout:   100 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Backend) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
