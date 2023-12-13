package main

import (
	"context"

	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/YourSurpriseCom/gcp-pubsub-deadletter-logger/config"
	"github.com/YourSurpriseCom/gcp-pubsub-deadletter-logger/server"
	log "github.com/jlentink/yaglogger"
)

func init() {
	log.SetLevel(log.LevelDebug)

	switch strings.ToLower(config.AppConfig.LogLevel) {
	case "info":
		log.SetLevel(log.LevelInfo)
	case "debug":
		log.SetLevel(log.LevelDebug)
	case "warning":
		log.SetLevel(log.LevelWarn)
	case "fatal":
		log.SetLevel(log.LevelFatal)
	default:
		log.SetLevel(log.LevelInfo)
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	server := server.New(config.AppConfig.Port)

	server.ListenOrExit()
	defer stopServer(ctx, server)

	<-getSignalChannel()

	log.Info("Caught signal, shutting down gracefully")
}

func getSignalChannel() chan os.Signal {
	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	return sigs
}

func stopServer(ctx context.Context, server server.Server) {
	err := server.Stop(ctx, 10*time.Second)
	if err != nil {
		log.Fatalf("Failed to stop server: %v", err)
	}
}
