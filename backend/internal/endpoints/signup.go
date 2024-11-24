package endpoints

import (
	"go-chat/internal/database"
	"go-chat/internal/dto"
	"go-chat/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SignUp(ctx *gin.Context) {
	var newUser dto.NewUser

	if err := ctx.ShouldBindJSON(&newUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	service := services.AuthService{
		Db: database.NewDb(),
	}

	token, err := service.SignUp(newUser)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"token": token})
}
