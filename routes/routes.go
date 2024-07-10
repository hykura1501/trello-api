package routes

import (
	"trello-api/handlers"

	"github.com/labstack/echo/v4"
)

type API struct {
	Echo        *echo.Echo
	UserHandler handlers.UserHandler
}

func (api *API) SetupRouter() {
	UserRouters(api.Echo, api.UserHandler)
}
