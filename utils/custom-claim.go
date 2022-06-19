package utils

import (
	"github.com/dgrijalva/jwt-go"
)

// CustomClaims custom jwt standard claims
type CustomClaims struct {
	jwt.StandardClaims
	Email          string   `json:"email"`
	Name           string   `json:"name"`
	LastName       string   `json:"last_name"`
	Permissions    []string `json:"permissions"`
	SignatureImage string   `json:"signature_image"`
}
