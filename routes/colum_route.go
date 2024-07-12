package routes

import (
	"trello-api/handlers"

	"github.com/labstack/echo/v4"
)

func ColumnRouters(c *echo.Echo, handler handlers.ColumnHandler) {
	c.POST("/column/new", handler.NewColumn)
}
