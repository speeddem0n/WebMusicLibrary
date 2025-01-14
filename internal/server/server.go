package server

import (
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

// Метод запуска сервера
func (s *Server) Run(host, port string, handler http.Handler) error {
	// Объявляем структуру http.Server
	s.httpServer = &http.Server{
		Addr:           host + ":" + port, // Server address
		Handler:        handler,           // Handler
		MaxHeaderBytes: 1 << 20,           // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

// Метод остановки сервера
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
