package routes

import (
	"trello-api/handlers"

	"github.com/labstack/echo/v4"
)

func CardRouters(c *echo.Echo, handler handlers.CardHandler) {
	c.POST("/card/new", handler.NewCard)
	c.GET("/card/detail/:card_id", handler.CardDetail)
	c.GET("/card/:column_id", handler.GetAllCards)
	c.PATCH("/card/update", handler.UpdateCard)
	c.POST("/card/attachment/new", handler.NewAttachment)
	c.POST("/card/attachment", handler.GetAllAttachments)
	c.DELETE("/card/attachment/delete", handler.DeleteAttachment)
}
