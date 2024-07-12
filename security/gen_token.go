package security

import (
	"os"
	"time"
	"trello-api/models"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user models.User) (string, time.Time, error) {
	var JWT_TOKEN_EXPIRES_AT = time.Now().Add(time.Hour * 24)
	claims := &models.JWTCustomsClaims{
		UserId: user.UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(JWT_TOKEN_EXPIRES_AT),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", JWT_TOKEN_EXPIRES_AT, err
	}
	return token, JWT_TOKEN_EXPIRES_AT, nil
}
