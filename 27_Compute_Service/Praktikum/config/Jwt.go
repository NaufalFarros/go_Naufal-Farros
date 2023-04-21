package config

import (
	"os"
	"github.com/golang-jwt/jwt/v5"
)

var JWT_KEY = []byte(os.Getenv("JWT_KEY"))

type JTWClaim struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
