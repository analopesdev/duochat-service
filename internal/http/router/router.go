package httpx

import (
	"net/http"

	"github.com/analopesdev/duochat-service/internal/user"
)

type RouterDeps struct {
	UserHandlers *user.Handler
}

func NewServer(addr string, deps RouterDeps) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", deps.UserHandlers.CreateUser)
	mux.HandleFunc("GET /users/", deps.UserHandlers.GetUserByID)
	mux.HandleFunc("GET /users/by-nickname/", deps.UserHandlers.GetUserByNickname)

	return &http.Server{
		Addr:    addr,
		Handler: mux,
	}
}
