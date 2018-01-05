package models

type User struct {
	Email    string `json:"email" yaml:"email"`
	Password string `json:"password" yaml:"password"`
	Base
}

func NewUser() User {
	return User{}
}
