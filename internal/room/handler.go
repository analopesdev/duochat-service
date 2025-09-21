package room

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

type CreateRoomRequestDTO struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	IsPrivate   bool    `json:"is_private"`
	Password    *string `json:"password"`
	MaxUsers    int     `json:"max_users"`
	CreatedBy   string  `json:"created_by"`
}

func (h *Handler) CreateRoom(w http.ResponseWriter, r *http.Request) {
	var dto CreateRoomRequestDTO

	if err := json.NewDecoder(r.Body).Decode(&dto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err := h.service.Create(r.Context(), &Room{
		Name:        dto.Name,
		Description: dto.Description,
		IsPrivate:   dto.IsPrivate,
		Password:    dto.Password,
		MaxUsers:    dto.MaxUsers,
		CreatedBy:   uuid.MustParse(dto.CreatedBy),
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Room created successfully"})
}

func (h *Handler) FindAllRooms(w http.ResponseWriter, r *http.Request) {
	rooms, err := h.service.FindAll(r.Context())

	print("FindAllRooms ============================")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(rooms)
}

func (h *Handler) GetRoomByID(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	idStr := strings.TrimPrefix(path, "/rooms/")

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}
}

func (h *Handler) DeleteRoom(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path

	idStr := strings.TrimPrefix(path, "/rooms/")
	createdByStr := r.URL.Query().Get("created_by")

	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	createdBy, err := uuid.Parse(createdByStr)
	if err != nil {
		http.Error(w, "invalid created_by", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(r.Context(), id, createdBy)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Room deleted successfully"})

}
