package routes

import (
	"trello-api/handlers"

	"github.com/labstack/echo/v4"
)

func UserRouters(e *echo.Echo, handler handlers.UserHandler) {
	e.GET("/user/login", handler.Login)
}
