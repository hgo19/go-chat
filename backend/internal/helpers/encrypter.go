package helpers

import (
	"log"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	saltRounds, err := strconv.Atoi(os.Getenv("SALT_ROUNDS"))
	if err != nil {
		log.Println("Error to get .env salt_rounds")
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), saltRounds)

	return string(bytes), err
}
