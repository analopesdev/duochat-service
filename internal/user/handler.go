package user

import (
	"net/http"
	"strconv"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	h.service.Create(r.Context(), &User{
		Nickname: r.FormValue("nickname"),
	})
}

func (h *Handler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(r.FormValue("id"), 10, 64)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	h.service.GetByID(r.Context(), id)
}

func (h *Handler) GetUserByNickname(w http.ResponseWriter, r *http.Request) {
	nickname := r.FormValue("nickname")

	if nickname == "" {
		http.Error(w, "nickname is required", http.StatusBadRequest)
		return
	}
	h.service.GetByNickname(r.Context(), nickname)
}
