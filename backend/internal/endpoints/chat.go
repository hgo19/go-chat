package endpoints

import (
	"go-chat/internal/helpers"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func ServerWs(c *gin.Context) {
	upgrader := websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		log.Printf("Error in upgrading web socket. Error: %v", err)
		return
	}

	go helpers.HandleClient(conn)
}
