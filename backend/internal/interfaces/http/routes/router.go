package routes

import (
	"log/slog"
	"net/http"

	"github.com/fairy-castle/family-story-universe/backend/internal/interfaces/http/handlers"
	"github.com/fairy-castle/family-story-universe/backend/internal/interfaces/http/middlewares"
)

// NewRouter wires HTTP routes for the Fairy Castle API.
func NewRouter(log *slog.Logger, version string) http.Handler {
	mux := http.NewServeMux()
	healthHandler := handlers.NewHealthHandler(version)

	mux.HandleFunc("GET /healthz", healthHandler.Health)
	mux.HandleFunc("GET /api/v1/healthz", healthHandler.Health)

	log.Info("router initialized", "version", version)
	return middlewares.RequestID(mux)
}
