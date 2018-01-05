package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Email    string
	Password string
	gorm.Model
}

func NewUser() User {
	return User{}
}
