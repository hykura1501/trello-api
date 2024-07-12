package routes

import (
	"trello-api/handlers"
	"trello-api/middlewares"

	"github.com/labstack/echo/v4"
)

func UserRouters(e *echo.Echo, handler handlers.UserHandler) {
	e.POST("/user/register", handler.Register)
	e.POST("/user/login", handler.Login)
	e.GET("/user/profile", handler.Profile, middlewares.JWTCustomsMiddleware())
}
