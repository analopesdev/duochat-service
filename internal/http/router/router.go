package httpx

import (
	"net/http"

	"github.com/analopesdev/duochat-service/internal/auth"
	"github.com/analopesdev/duochat-service/internal/http/middleware"
	"github.com/analopesdev/duochat-service/internal/room"
	"github.com/analopesdev/duochat-service/internal/user"
	"github.com/analopesdev/duochat-service/internal/ws"
)

type RouterDeps struct {
	UserHandlers *user.Handler
	RoomHandlers *room.Handler
	WsHandler    *ws.Handler
	WsHub        *ws.Hub
	AuthHandler  *auth.Handler
}

func NewServer(addr string, deps RouterDeps) *http.Server {
	mux := http.NewServeMux()

	middlewareAuth := middleware.AuthMiddleware()

	middleware.HandleFunc(mux, "POST /auth", deps.AuthHandler.AuthUser)
	middleware.HandleFunc(mux, "POST /users", deps.UserHandlers.CreateUser)

	middleware.HandleFunc(mux, "GET /users", deps.UserHandlers.FindAllUsers, middlewareAuth)
	middleware.HandleFunc(mux, "GET /users/{id}", deps.UserHandlers.GetUserByID, middlewareAuth)
	middleware.HandleFunc(mux, "GET /users/by-nickname/{nickname}", deps.UserHandlers.GetUserByNickname, middlewareAuth)

	middleware.HandleFunc(mux, "POST /rooms", deps.RoomHandlers.CreateRoom, middlewareAuth)
	middleware.HandleFunc(mux, "GET /rooms", deps.RoomHandlers.FindAllRooms, middlewareAuth)
	middleware.HandleFunc(mux, "GET /rooms/{id}", deps.RoomHandlers.GetRoomByID, middlewareAuth)
	middleware.HandleFunc(mux, "DELETE /rooms/{id}", deps.RoomHandlers.DeleteRoom, middlewareAuth)

	middleware.HandleFunc(mux, "GET /ws", deps.WsHandler.ServeWs)

	corsOptions := middleware.DefaultCORSOptions()
	handler := middleware.CORS(corsOptions)(mux)

	return &http.Server{
		Addr:    addr,
		Handler: handler,
	}
}
