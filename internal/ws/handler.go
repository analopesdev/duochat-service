package ws

import (
	"log"
	"net/http"
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
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &Client{hub: h.hub, conn: conn, send: make(chan []byte, 256)}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()
}
