package routes

import (
	"trello-api/handlers"
	"trello-api/middlewares"

	"github.com/labstack/echo/v4"
)

func ColumnRouters(c *echo.Echo, handler handlers.ColumnHandler) {
	columnGroup := c.Group("/column", middlewares.JWTCustomsMiddleware())
	{
		columnGroup.POST("/new", handler.NewColumn)
		columnGroup.GET("/detail/:column_id", handler.ColumnDetail)
		columnGroup.GET("/:board_id", handler.GetAllColumns)
		columnGroup.PATCH("/update", handler.UpdateColumn)
	}

}
