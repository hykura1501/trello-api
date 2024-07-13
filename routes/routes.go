package routes

import (
	"trello-api/handlers"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo          *echo.Echo
	UserHandler   handlers.UserHandler
	BoardHandler  handlers.BoardHandler
	ColumnHandler handlers.ColumnHandler
	CardHandler   handlers.CardHandler
}

func (api *API) SetupRouter() {
	UserRouters(api.Echo, api.UserHandler)
	BoardRouters(api.Echo, api.BoardHandler)
	ColumnRouters(api.Echo, api.ColumnHandler)
	CardRouters(api.Echo, api.CardHandler)
}
