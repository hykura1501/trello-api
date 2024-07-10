package middlewares

import (
	"os"
	"trello-api/models"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"

	echojwt "github.com/labstack/echo-jwt/v4"
)

func JWTCustomsMiddleware() echo.MiddlewareFunc {
	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(models.JWTCustomsClaims)
		},
		SigningKey: []byte(os.Getenv("JWT_SECRET_KEY")),
	}
	return echojwt.WithConfig(config)
}
