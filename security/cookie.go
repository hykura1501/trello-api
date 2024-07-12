package security

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func WriteCookie(c echo.Context, name, value string, expiresAt time.Time) {
	cookie := new(http.Cookie)
	cookie.Name = name
	cookie.Value = value
	cookie.Expires = expiresAt
	cookie.Path = "/"
	c.SetCookie(cookie)
}

func ReadCookie(c echo.Context, name string) (string, error) {
	cookie, err := c.Cookie(name)
	if err != nil {
		return "", err
	}
	return cookie.Name, nil
}
