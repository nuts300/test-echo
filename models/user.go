package models

import (
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Email    string `json:"email" yaml:"email" validate:"required,email"`
	Password string `json:"password" yaml:"password" validate:"required"`
	Base
}

func (user *User) BeforeSave(scope *gorm.Scope) (err error) {
	if pw, err := bcrypt.GenerateFromPassword([]byte(user.Password), 0); err == nil {
		scope.SetColumn("password", pw)
		return err
	}
	return nil
}

func NewUser() User {
	return User{}
}
