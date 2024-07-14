package routes

import (
	"trello-api/handlers"

	"github.com/labstack/echo/v4"
)

func ColumnRouters(c *echo.Echo, handler handlers.ColumnHandler) {
	c.POST("/column/new", handler.NewColumn)
	c.GET("/column/detail/:column_id", handler.ColumnDetail)
	c.GET("/column/:board_id", handler.GetAllColumns)
	c.PATCH("/column/update", handler.UpdateColumn)
}
