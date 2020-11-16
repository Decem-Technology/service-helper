package bootstrap

import (
	"crypto/rsa"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/Decem-Technology/service-helper/utils"
	"github.com/dgrijalva/jwt-go"
)

type (
	// OAuth oauth helper
	OAuth struct {
	}

	tokenType int

	grantType string
)

var (
	// PrivateKey rsa private key
	PrivateKey *rsa.PrivateKey
	// PublicKey rsa public key
	PublicKey *rsa.PublicKey
	// ErrInvalidToken set error valid token
	ErrInvalidToken = errors.New("token: invalid token")
	// AccessTokenDuration expire token time
	AccessTokenDuration = time.Duration(time.Hour * 24 * 14)
	// RefreshTokenDuration expire token time
	RefreshTokenDuration = time.Duration(time.Hour * 24 * 20)
)

// Grant Type
const (
	_                               = iota
	TokenTypeRefreshToken tokenType = iota
	TokenTypeAccessToken
	GrantTypePassword     = "password"
	GrantTypeRefreshToken = "refresh_token"
)

func init() {
	pwd, _ := os.Getwd()
	if os.Getenv("LOAD_PRIVATE_KEY") == "true" {
		key, err := ioutil.ReadFile(pwd + "/storage/private.key")
		if err != nil {
			panic(fmt.Sprintf("[OAUTH] private.key %s", err))
		}
		PrivateKey, err = jwt.ParseRSAPrivateKeyFromPEM(key)
		if err != nil {
			panic(fmt.Sprintf("[OAUTH] Could not parse private.key error: %s", err))
		}
	}
	if os.Getenv("LOAD_PUBLIC_KEY") == "true" {
		key, err := ioutil.ReadFile(pwd + "/storage/public.key")
		if err != nil {
			panic(fmt.Sprintf("[OAUTH] public.key %s", err))
		}
		PublicKey, err = jwt.ParseRSAPublicKeyFromPEM(key)
		if err != nil {
			panic(fmt.Sprintf("[OAUTH] Could not parse public.key error: %s", err))
		}
	}
}

// VerifyJWT verify jwt token
func (ctl OAuth) VerifyJWT(t string) (*utils.CustomClaims, error) {
	tok, err := jwt.ParseWithClaims(t, &utils.CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Check is token use correct signing method
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %s", token.Header["alg"])
		}
		// return secret for this signing method
		return PublicKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tok.Claims.(*utils.CustomClaims); ok && tok.Valid {
		return claims, nil
	}
	return nil, errors.New("token: invalid token")
}
