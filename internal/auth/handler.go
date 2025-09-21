package auth

import (
	"encoding/json"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: *service}
}

func (h *Handler) AuthUser(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")

	if token == "" {
		http.Error(w, "token is required", http.StatusBadRequest)
		return
	}

	valid := h.service.ValidateToken(token)

	if !valid {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User authenticated successfully"})
}
