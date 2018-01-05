package models

type User struct {
	Email    string `json:"email" yaml:"email" validate:"required,email"`
	Password string `json:"password" yaml:"password" validate:"required"`
	Base
}

func NewUser() User {
	return User{}
}
