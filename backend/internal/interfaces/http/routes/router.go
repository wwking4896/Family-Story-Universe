package routes

import (
	"log/slog"
	"net/http"

	"github.com/fairy-castle/family-story-universe/backend/internal/application/services"
	"github.com/fairy-castle/family-story-universe/backend/internal/interfaces/http/handlers"
	"github.com/fairy-castle/family-story-universe/backend/internal/interfaces/http/middlewares"
)

// NewRouter wires HTTP routes for the Fairy Castle API.
func NewRouter(log *slog.Logger, version string, store *services.MemoryStore) http.Handler {
	mux := http.NewServeMux()
	healthHandler := handlers.NewHealthHandler(version)
	apiHandler := handlers.NewAPIHandler(store)

	mux.HandleFunc("GET /healthz", healthHandler.Health)
	mux.HandleFunc("GET /api/v1/healthz", healthHandler.Health)

	mux.HandleFunc("POST /api/v1/auth/register", apiHandler.Register)
	mux.HandleFunc("POST /api/v1/auth/login", apiHandler.Login)
	mux.HandleFunc("POST /api/v1/auth/logout", apiHandler.Logout)
	mux.HandleFunc("GET /api/v1/auth/me", apiHandler.Me)
	mux.HandleFunc("POST /api/v1/families", apiHandler.CreateFamily)
	mux.HandleFunc("GET /api/v1/families/me", apiHandler.MyFamilies)
	mux.HandleFunc("PATCH /api/v1/families/{familyId}", apiHandler.UpdateFamily)
	mux.HandleFunc("POST /api/v1/children", apiHandler.CreateChild)
	mux.HandleFunc("GET /api/v1/children", apiHandler.ListChildren)
	mux.HandleFunc("GET /api/v1/children/{childId}", apiHandler.GetChild)
	mux.HandleFunc("PATCH /api/v1/children/{childId}", apiHandler.UpdateChild)
	mux.HandleFunc("DELETE /api/v1/children/{childId}", apiHandler.DeleteChild)
	mux.HandleFunc("POST /api/v1/characters", apiHandler.CreateCharacter)
	mux.HandleFunc("GET /api/v1/characters", apiHandler.ListCharacters)
	mux.HandleFunc("GET /api/v1/characters/{characterId}", apiHandler.GetCharacter)
	mux.HandleFunc("PATCH /api/v1/characters/{characterId}", apiHandler.UpdateCharacter)
	mux.HandleFunc("DELETE /api/v1/characters/{characterId}", apiHandler.DeleteCharacter)
	mux.HandleFunc("GET /api/v1/regions", apiHandler.Regions)
	mux.HandleFunc("POST /api/v1/stories/generate", apiHandler.GenerateStory)
	mux.HandleFunc("GET /api/v1/stories", apiHandler.ListStories)
	mux.HandleFunc("GET /api/v1/stories/{storyId}", apiHandler.GetStory)
	mux.HandleFunc("PATCH /api/v1/stories/{storyId}", apiHandler.UpdateStory)
	mux.HandleFunc("DELETE /api/v1/stories/{storyId}", apiHandler.DeleteStory)
	mux.HandleFunc("GET /api/v1/timebook", apiHandler.Timebook)
	mux.HandleFunc("GET /api/v1/timebook/{year}", apiHandler.TimebookYear)

	log.Info("router initialized", "version", version)
	return middlewares.RequestID(mux)
}
