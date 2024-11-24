package services

import (
	"go-chat/internal/dto"
	"go-chat/internal/helpers"
	"go-chat/internal/models"

	"gorm.io/gorm"
)

type AuthService struct {
	Db *gorm.DB
}

func (s *AuthService) SignUp(newUser dto.NewUser) (string, error) {
	hashedPassword, err := helpers.HashPassword(newUser.Password)
	if err != nil {
		return "", err
	}

	createdUser := models.NewUser(newUser.Username, hashedPassword)

	tx := s.Db.Create(&createdUser)

	token, err := helpers.CreateToken(createdUser.Username)

	if err != nil {
		return "", err
	}

	return token, tx.Error
}
