package main

import (
	"go-chat/internal/endpoints"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	router.GET("/ws", endpoints.ServerWs)
	err := router.Run(":3001")
	if err != nil {
		log.Fatalf("Unable to start server. Error %v", err)
	}
	log.Println("Server started successfully.")
}
