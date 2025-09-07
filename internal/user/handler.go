package user

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

type CreateUserRequestDTO struct {
	Nickname string `json:"nickname"`
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var body struct {
		Nickname string `json:"nickname"`
	}
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	token, err := h.service.Create(r.Context(), &User{Nickname: body.Nickname})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{Name: "uid", Value: token, HttpOnly: true, SameSite: http.SameSiteLaxMode, MaxAge: 30 * 24 * 3600})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

func (h *Handler) FindAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.service.FindAll(r.Context())

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"users":   users,
			"message": "Users fetched successfully",
		},
	)
}

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	idStr := strings.TrimPrefix(path, "/users/")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	u, err := h.service.GetByID(r.Context(), id)
	if err != nil {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		map[string]interface{}{
			"user":    u,
			"message": "User fetched successfully",
		},
	)
}

func (h *Handler) GetUserByNickname(w http.ResponseWriter, r *http.Request) {
	nickname := r.FormValue("nickname")

	if nickname == "" {
		http.Error(w, "nickname is required", http.StatusBadRequest)
		return
	}
	h.service.GetByNickname(r.Context(), nickname)
}
