package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Email    string `json:"email" yaml:"email"`
	Password string `json:"password" yaml:"password"`
	gorm.Model
}

func NewUser() User {
	return User{}
}
