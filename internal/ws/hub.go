package ws

import (
	"log"
)

type Hub struct {
	rooms map[string]map[*Client]bool

	register   chan *Client
	unregister chan *Client
	broadcast  chan *Message
}

type Message struct {
	RoomID  string `json:"room_id"`
	Content []byte `json:"content"`
	UserID  int64  `json:"user_id"`
}

func newHub() *Hub {
	return &Hub{
		rooms:      make(map[string]map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan *Message),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.registerClient(client)

		case client := <-h.unregister:
			h.unregisterClient(client)

		case message := <-h.broadcast:
			h.broadcastToRoom(message)
		}
	}
}

func (h *Hub) registerClient(client *Client) {
	// Se a sala não existe, cria ela
	if h.rooms[client.RoomID] == nil {
		h.rooms[client.RoomID] = make(map[*Client]bool)
	}

	// Adiciona o cliente à sala
	h.rooms[client.RoomID][client] = true

	log.Printf("Cliente %d entrou na sala %s", client.ID, client.RoomID)
}

func (h *Hub) unregisterClient(client *Client) {
	if room, exists := h.rooms[client.RoomID]; exists {
		if _, clientExists := room[client]; clientExists {
			delete(room, client)
			close(client.send)

			// Se a sala ficou vazia, remove ela
			if len(room) == 0 {
				delete(h.rooms, client.RoomID)
			}

			log.Printf("Cliente %d saiu da sala %s", client.ID, client.RoomID)
		}
	}
}

func (h *Hub) broadcastToRoom(message *Message) {
	if room, exists := h.rooms[message.RoomID]; exists {
		for client := range room {
			select {
			case client.send <- message.Content:
			default:
				close(client.send)
				delete(room, client)
			}
		}
	}
}
