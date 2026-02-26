package infraestructure

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		// En producción, validar el origen adecuadamente
		return true
	},
}

// HandleWebSocket maneja nuevas conexiones WebSocket
func HandleWebSocket(hub *Hub) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}

		client := &Client{
			hub:     hub,
			conn:    conn,
			send:    make(chan WebSocketMessage, 256),
			grupoID: 0,
		}

		hub.register <- client

		go client.writePump()
		go client.readPump()
	}
}
