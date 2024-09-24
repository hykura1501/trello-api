package routes

import (
	"trello-api/handlers"
	"trello-api/middlewares"

	"github.com/labstack/echo/v4"
)

func BoardRouters(c *echo.Echo, handler handlers.BoardHandler) {
	boardGroup := c.Group("/board", middlewares.JWTCustomsMiddleware())
	{
		boardGroup.POST("/new", handler.NewBoard)
		boardGroup.GET("/detail/:board_id", handler.BoardDetail)
	}
}
