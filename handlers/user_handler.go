package handlers

import (
	"net/http"
	"trello-api/repository"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserRepo repository.UserRepository
}

func (handler UserHandler) Login(c echo.Context) error {
	return c.JSON(http.StatusOK, "Login")
}
