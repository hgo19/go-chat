package helpers

import (
	"go-chat/internal/models"
	"log"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func broadcast(msg models.Message) {
	for conn := range clients {
		conn.WriteJSON(msg)
	}
}

var clients = make(map[*websocket.Conn]struct{})

func HandleClient(c *websocket.Conn) {
	defer func() {
		delete(clients, c)
		log.Println("Closing Websocket")
		c.Close()
	}()

	clients[c] = struct{}{}

	for {
		var msg models.Message

		err := c.ReadJSON(&msg)
		if err != nil {
			log.Printf("Error in reading json message. Error : %v", err)
			return
		}

		broadcast(msg)
	}
}
