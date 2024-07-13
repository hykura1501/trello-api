package routes

import (
	"trello-api/handlers"

	"github.com/labstack/echo/v4"
)

func BoardRouters(c *echo.Echo, handler handlers.BoardHandler) {
	c.POST("/board/new", handler.NewBoard)
	c.GET("/board/detail/:board_id", handler.BoardDetail)
}
