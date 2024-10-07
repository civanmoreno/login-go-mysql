package auth

import "github.com/golang-jwt/jwt/v4"

type Claims struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}
