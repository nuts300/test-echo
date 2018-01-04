package userResource

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Email    string
	Password string
	gorm.Model
}

func New() User {
	return User{}
}
