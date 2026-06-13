package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

// HealthHandler exposes service health endpoints.
type HealthHandler struct {
	startedAt time.Time
	version   string
}

// NewHealthHandler creates a health handler.
func NewHealthHandler(version string) *HealthHandler {
	return &HealthHandler{startedAt: time.Now().UTC(), version: version}
}

// Health returns a lightweight liveness response.
func (h *HealthHandler) Health(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(map[string]string{
		"status":     "ok",
		"service":    "fairy-castle-api",
		"version":    h.version,
		"started_at": h.startedAt.Format(time.RFC3339),
	})
}
