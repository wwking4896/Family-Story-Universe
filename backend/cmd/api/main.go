package main

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fairy-castle/family-story-universe/backend/internal/config"
	"github.com/fairy-castle/family-story-universe/backend/internal/interfaces/http/routes"
	"github.com/fairy-castle/family-story-universe/backend/pkg/logger"
)

const version = "0.1.0"

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel)
	slog.SetDefault(log)

	router := routes.NewRouter(log, version)
	server := &http.Server{
		Addr:              ":" + cfg.HTTPPort,
		Handler:           router,
		ReadHeaderTimeout: 5 * time.Second,
	}

	go func() {
		log.Info("starting Fairy Castle API", "addr", server.Addr, "env", cfg.AppEnv)
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("api server failed", "error", err)
			os.Exit(1)
		}
	}()

	shutdown(log, server)
}

func shutdown(log *slog.Logger, server *http.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	log.Info("shutting down Fairy Castle API")
	if err := server.Shutdown(ctx); err != nil {
		log.Error("forced shutdown", "error", err)
	}
}
