package handlers

import (
	"encoding/json"
	storage "gahmen-api/db"
	"net/http"
)

type Handler struct {
	store storage.Storage
}

func NewHandler(store storage.Storage) *Handler {
	return &Handler{store: store}
}

// @Summary Health Check
// @Description Check if the API is up and running
// @Tags Health
// @Produce json
// @Success 200 {object} map[string]string "status: UP"
// @Router /health [get]
func (h *Handler) HealthCheck(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(map[string]string{"status": "UP"})
}

func (h *Handler) NotFound(w http.ResponseWriter, r *http.Request) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	return json.NewEncoder(w).Encode(map[string]string{"error": "Not Found"})
}
