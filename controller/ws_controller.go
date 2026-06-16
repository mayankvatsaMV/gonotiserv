// controller/ws_controller.go

package controller

import (
	"gonotiserv/websocket"
	"log"
	"net/http"

	gws "github.com/gorilla/websocket"

	"github.com/gin-gonic/gin"
)

var upgrader = gws.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WebSocketHandler(
	hub *websocket.Hub,
) gin.HandlerFunc {

	return func(c *gin.Context) {

		conn, err := upgrader.Upgrade(
			c.Writer,
			c.Request,
			nil,
		)

		if err != nil {
			return
		}

		hub.Mu.Lock()
		hub.Clients[conn] = true
		hub.Mu.Unlock()

		log.Println("client connected")

		defer func() {
			hub.Mu.Lock()
			delete(hub.Clients, conn)
			hub.Mu.Unlock()

			conn.Close()

			log.Println("client disconnected")
		}()

		for {
			_, _, err := conn.ReadMessage()

			if err != nil {
				break
			}
		}
	}
}
