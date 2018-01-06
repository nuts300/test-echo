package models

type AuthInfo struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	Token string `json:"string"`
}
