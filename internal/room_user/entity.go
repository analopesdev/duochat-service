package room_user

import (
	"time"

	"github.com/google/uuid"
)

type RoomUser struct {
	ID       uuid.UUID `json:"id" db:"id"`
	RoomID   uuid.UUID `json:"room_id" db:"room_id"`
	UserID   uuid.UUID `json:"user_id" db:"user_id"`
	IsAdmin  bool      `json:"is_admin" db:"is_admin"`
	JoinedAt time.Time `json:"joined_at" db:"joined_at"`
	IsActive bool      `json:"is_active" db:"is_active"`
}

func NewRoomUser(roomID uuid.UUID, userID uuid.UUID, isAdmin bool) *RoomUser {
	return &RoomUser{
		ID:       uuid.New(),
		RoomID:   roomID,
		UserID:   userID,
		IsAdmin:  isAdmin,
		JoinedAt: time.Now(),
		IsActive: true,
	}
}
