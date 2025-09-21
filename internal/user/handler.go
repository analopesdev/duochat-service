package user

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/analopesdev/duochat-service/internal/auth"
	"github.com/google/uuid"
)

type Handler struct {
	service     *Service
	authService *auth.Service
}

func NewHandler(service *Service, authService *auth.Service) *Handler {
	return &Handler{service: service, authService: authService}
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

	user, err := h.service.Create(r.Context(), &User{Nickname: body.Nickname})

	switch err {
	case ErrConflict:
		errorMsg, _ := json.Marshal(map[string]string{
			"message": "nickname already exists",
			"error":   err.Error(),
		})

		http.Error(w, string(errorMsg), http.StatusConflict)
		return
	case nil:

		token, err := h.authService.GenerateJWT(user.ID.String(), user.Nickname)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{Name: "uid", Value: token, HttpOnly: true, SameSite: http.SameSiteLaxMode, MaxAge: 30 * 24 * 3600})

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})

	default:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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

	u, err := h.service.GetByID(r.Context(), uuid.MustParse(idStr))
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
