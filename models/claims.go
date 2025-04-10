package models

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	Id   uint `json:"id`
	Role uint `json"role"`
	jwt.RegisteredClaims
}
