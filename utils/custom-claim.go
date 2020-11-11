package utils

import (
	"github.com/dgrijalva/jwt-go"
)

// CustomClaims custom jwt standard claims
type CustomClaims struct {
	jwt.StandardClaims
	Permissions []string `json:"permissions"`
}
