package models

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Username string `json:"username`
	Role     uint   `json"role"`
	jwt.RegisteredClaims
}
