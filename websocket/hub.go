// websocket/hub.go

package websocket

import (
	"log"
	"sync"

	"github.com/gorilla/websocket"
)

type Hub struct {
	Clients map[*websocket.Conn]bool
	Mu      sync.RWMutex
}

func NewHub() *Hub {
	return &Hub{
		Clients: make(map[*websocket.Conn]bool),
	}
}

func (h *Hub) Broadcast(msg any) {

	h.Mu.RLock()
	defer h.Mu.RUnlock() //defer

	for conn := range h.Clients {

		err := conn.WriteJSON(msg)

		if err != nil {
			log.Println(err)
		}
	}
}
