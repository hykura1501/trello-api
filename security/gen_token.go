package security

import (
	"os"
	"time"
	"trello-api/models"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateToken(user models.User) (string, error) {
	claims := &models.JWTCustomsClaims{
		UserId: user.UserId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return token, nil
}
