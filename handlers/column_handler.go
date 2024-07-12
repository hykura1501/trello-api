package handlers

import (
	"trello-api/repository"

	"github.com/labstack/echo/v4"
)

type ColumnHandler struct {
	ColumnRepo repository.ColumnRepository
}

func (repo ColumnHandler) NewColumn(c echo.Context) error {
	return nil
}
