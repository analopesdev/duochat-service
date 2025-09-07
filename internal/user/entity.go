package user

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID           uuid.UUID `json:"id"`
	Nickname     string    `json:"nickname"`
	Avatar       string    `json:"avatar"`
	TokenVersion int       `json:"token_version"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func NewUser(nickname string) *User {
	return &User{
		ID:           uuid.New(),
		Nickname:     nickname,
		Avatar:       "https://ui-avatars.com/api/?name=" + nickname,
		TokenVersion: 1,
		UpdatedAt:    time.Now(),
	}
}
