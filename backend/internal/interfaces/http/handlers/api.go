package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/fairy-castle/family-story-universe/backend/internal/application/services"
	"github.com/fairy-castle/family-story-universe/backend/internal/domain"
)

type APIHandler struct {
	store *services.MemoryStore
}

func NewAPIHandler(store *services.MemoryStore) *APIHandler {
	return &APIHandler{store: store}
}

type registerRequest struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	DisplayName string `json:"display_name"`
}

type loginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type familyRequest struct {
	Name string `json:"name"`
}

func (h *APIHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req registerRequest
	if !decodeJSON(w, r, &req) {
		return
	}
	result, err := h.store.Register(req.Email, req.Password, req.DisplayName)
	respond(w, result, err, http.StatusCreated)
}

func (h *APIHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req loginRequest
	if !decodeJSON(w, r, &req) {
		return
	}
	result, err := h.store.Login(req.Email, req.Password)
	respond(w, result, err, http.StatusOK)
}

func (h *APIHandler) Logout(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *APIHandler) Me(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"user": user, "families": h.store.MyFamilies(user.ID)})
}

func (h *APIHandler) CreateFamily(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	var req familyRequest
	if !decodeJSON(w, r, &req) {
		return
	}
	family, err := h.store.CreateFamily(user.ID, req.Name)
	respond(w, family, err, http.StatusCreated)
}

func (h *APIHandler) MyFamilies(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	writeJSON(w, http.StatusOK, map[string]any{"items": h.store.MyFamilies(user.ID)})
}

func (h *APIHandler) CreateChild(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	var child domain.Child
	if !decodeJSON(w, r, &child) {
		return
	}
	created, err := h.store.CreateChild(user.ID, child)
	respond(w, created, err, http.StatusCreated)
}

func (h *APIHandler) ListChildren(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	familyID, ok := intQuery(w, r, "family_id")
	if !ok {
		return
	}
	items, err := h.store.ListChildren(user.ID, familyID)
	respond(w, map[string]any{"items": items}, err, http.StatusOK)
}

func (h *APIHandler) CreateCharacter(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	var character domain.Character
	if !decodeJSON(w, r, &character) {
		return
	}
	created, err := h.store.CreateCharacter(user.ID, character)
	respond(w, created, err, http.StatusCreated)
}

func (h *APIHandler) ListCharacters(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	familyID, ok := intQuery(w, r, "family_id")
	if !ok {
		return
	}
	items, err := h.store.ListCharacters(user.ID, familyID)
	respond(w, map[string]any{"items": items}, err, http.StatusOK)
}

func (h *APIHandler) Regions(w http.ResponseWriter, _ *http.Request) {
	writeJSON(w, http.StatusOK, map[string]any{"items": h.store.Regions()})
}

func (h *APIHandler) GenerateStory(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	var input services.StoryGenerateInput
	if !decodeJSON(w, r, &input) {
		return
	}
	result, err := h.store.GenerateStory(user.ID, input)
	respond(w, result, err, http.StatusCreated)
}

func (h *APIHandler) ListStories(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	familyID, ok := intQuery(w, r, "family_id")
	if !ok {
		return
	}
	items, err := h.store.ListStories(user.ID, familyID)
	respond(w, map[string]any{"items": items}, err, http.StatusOK)
}

func (h *APIHandler) GetStory(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	storyID, err := strconv.ParseInt(r.PathValue("storyId"), 10, 64)
	if err != nil {
		writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", "storyId 格式不正確")
		return
	}
	story, err := h.store.GetStory(user.ID, storyID)
	respond(w, story, err, http.StatusOK)
}

func (h *APIHandler) Timebook(w http.ResponseWriter, r *http.Request) {
	user, ok := h.requireUser(w, r)
	if !ok {
		return
	}
	familyID, ok := intQuery(w, r, "family_id")
	if !ok {
		return
	}
	result, err := h.store.Timebook(user.ID, familyID)
	respond(w, result, err, http.StatusOK)
}

func (h *APIHandler) requireUser(w http.ResponseWriter, r *http.Request) (domain.User, bool) {
	token := strings.TrimPrefix(r.Header.Get("Authorization"), "Bearer ")
	if token == "" {
		writeError(w, http.StatusUnauthorized, "UNAUTHORIZED", "請先登入")
		return domain.User{}, false
	}
	user, err := h.store.UserFromToken(token)
	if err != nil {
		writeError(w, http.StatusUnauthorized, "UNAUTHORIZED", "登入狀態已失效")
		return domain.User{}, false
	}
	return user, true
}

func decodeJSON(w http.ResponseWriter, r *http.Request, dst any) bool {
	defer r.Body.Close()
	if err := json.NewDecoder(r.Body).Decode(dst); err != nil {
		writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", "JSON 格式不正確")
		return false
	}
	return true
}

func intQuery(w http.ResponseWriter, r *http.Request, key string) (int64, bool) {
	value := r.URL.Query().Get(key)
	id, err := strconv.ParseInt(value, 10, 64)
	if err != nil || id <= 0 {
		writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", key+" 為必填且必須是正整數")
		return 0, false
	}
	return id, true
}

func respond(w http.ResponseWriter, payload any, err error, successStatus int) {
	if err != nil {
		handleError(w, err)
		return
	}
	writeJSON(w, successStatus, payload)
}

func handleError(w http.ResponseWriter, err error) {
	switch {
	case errors.Is(err, services.ErrUnauthorized):
		writeError(w, http.StatusUnauthorized, "UNAUTHORIZED", "請先登入")
	case errors.Is(err, services.ErrForbidden):
		writeError(w, http.StatusForbidden, "FORBIDDEN", "沒有權限存取此資源")
	case errors.Is(err, services.ErrNotFound):
		writeError(w, http.StatusNotFound, "NOT_FOUND", "找不到資源")
	case errors.Is(err, services.ErrConflict):
		writeError(w, http.StatusConflict, "CONFLICT", "資料已存在")
	case errors.Is(err, services.ErrValidation):
		writeError(w, http.StatusBadRequest, "VALIDATION_ERROR", "欄位格式不正確")
	default:
		writeError(w, http.StatusInternalServerError, "INTERNAL_ERROR", "系統發生錯誤")
	}
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeError(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, map[string]any{"error": map[string]any{"code": code, "message": message, "details": map[string]any{}}})
}
