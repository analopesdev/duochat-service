package user

import (
	"time"
)

type User struct {
	ID        int64     `json:"id"`
	Nickname  string    `json:"nickname"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(nickname string) *User {
	return &User{
		Nickname:  nickname,
		UpdatedAt: time.Now(),
	}
}
