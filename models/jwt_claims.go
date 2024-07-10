package models

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTCustomsClaims struct {
	UserId string
	jwt.RegisteredClaims
}
