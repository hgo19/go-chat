package main

import (
	"go-chat/internal/endpoints"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	router.POST("/signup", endpoints.SignUp)
	router.GET("/ws", endpoints.ServerWs)

	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"data": "Go chat API"})
	})
	err = router.Run(":3001")
	if err != nil {
		log.Fatalf("Unable to start server. Error %v", err)
	}
	log.Println("Server started successfully.")
}
