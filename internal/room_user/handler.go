package room_user

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

type JoinRoomRequestDTO struct {
	UserID   string  `json:"user_id"`
	RoomID   string  `json:"room_id"`
	Password *string `json:"password"`
}
