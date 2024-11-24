package models

import "github.com/rs/xid"

type User struct {
	ID       string `json:"id" gorm:"primary_key"`
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
}

func NewUser(username string, password string) *User {
	user := &User{
		ID:       xid.New().String(),
		Username: username,
		Password: password,
	}

	return user
}
