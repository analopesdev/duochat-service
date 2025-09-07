package ws

import (
	"log"
	"net/http"
	"strconv"
)

type Handler struct {
	hub *Hub
}

func NewHandler() *Handler {
	hub := newHub()
	go hub.run()
	return &Handler{hub: hub}
}

func (h *Handler) ServeWs(w http.ResponseWriter, r *http.Request) {
	roomID := r.URL.Query().Get("room")
	userIDStr := r.URL.Query().Get("user_id")

	if roomID == "" {
		http.Error(w, "room parameter is required", http.StatusBadRequest)
		return
	}

	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		http.Error(w, "invalid user_id", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
		return
	}

	client := &Client{
		hub:    h.hub,
		conn:   conn,
		ID:     userID,
		RoomID: roomID,
		send:   make(chan []byte, 256),
	}

	log.Printf("Cliente %d conectando na sala %s", userID, roomID)
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}
