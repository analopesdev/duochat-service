package user

import "errors"

var (
	ErrInvalidNick = errors.New("invalid nickname")
	ErrConflict    = errors.New("nickname already in use")
	ErrNotFound    = errors.New("user not found")
)
