package model

import "github.com/dgrijalva/jwt-go"

type Auth struct {
	Email string   `json:"email"`
	Role  []string `json:"role"`
	jwt.StandardClaims
}

type AuthRefreshToken struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
