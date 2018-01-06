package models

import jwt "github.com/dgrijalva/jwt-go"

type Claims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}
