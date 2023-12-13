package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	log "github.com/jlentink/yaglogger"
)

type Server struct {
	*http.Server
}

func New(port int) Server {
	router := GetRouter()

	server := &http.Server{
		ReadHeaderTimeout: time.Second * 60,
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           router,
	}

	return Server{server}
}

func (server Server) ListenOrExit() {
	go func() {
		log.Info("Starting HTTP server listening on '%s'", server.Addr)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Could not start HTTP server: %v", err)
		}
	}()
}

func (server Server) Stop(ctx context.Context, timeout time.Duration) error {
	shutdownContext, shutdownContextRelease := context.WithTimeout(ctx, timeout)
	defer shutdownContextRelease()

	err := server.Shutdown(shutdownContext)
	if err != nil {
		return fmt.Errorf("shutdown of HTTP server failed: %w", err)
	}
	return nil
}
