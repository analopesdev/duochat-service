package auth

import "github.com/google/uuid"

type Claims struct {
	UserID   uuid.UUID `json:"user_id"`
	Nickname string    `json:"nickname"`
	Exp      int64     `json:"exp"`
	Iat      int64     `json:"iat"`
}
