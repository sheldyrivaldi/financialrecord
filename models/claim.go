package models

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	ID       int    `json:"id"`
	Role     string `json:"role"`
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}
