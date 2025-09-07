package room

import (
	"time"

	"github.com/google/uuid"
)

type Room struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	IsPrivate   bool      `json:"is_private" db:"is_private"`
	Password    *string   `json:"-" db:"password"`
	MaxUsers    int       `json:"max_users" db:"max_users"`
	CreatedBy   uuid.UUID `json:"created_by" db:"created_by"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

func NewRoom(name string, description string, isPrivate bool, password *string, maxUsers int, createdBy uuid.UUID) *Room {
	return &Room{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
		IsPrivate:   isPrivate,
		Password:    password,
		MaxUsers:    maxUsers,
		CreatedBy:   createdBy,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
