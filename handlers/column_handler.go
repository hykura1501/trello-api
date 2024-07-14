package handlers

import (
	"log"
	"net/http"
	"time"
	"trello-api/models"
	"trello-api/repository"

	validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ColumnHandler struct {
	ColumnRepo repository.ColumnRepository
}

func (repo ColumnHandler) NewColumn(c echo.Context) error {
	var column models.Column
	if err := c.Bind(&column); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	validate := validator.New()
	if err := validate.Struct(column); err != nil {
		return c.JSON(http.StatusNotAcceptable, models.Response{
			Code:    http.StatusNotAcceptable,
			Message: err.Error(),
		})
	}
	// insert to db
	column.ColumnId = uuid.New().String()
	column.CreatedAt = time.Now()
	column.UpdatedAt = time.Now()
	if err := repo.ColumnRepo.SaveColumn(&column); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "column created",
		Data:    column,
	})
}

// [GET] /column/detail/:column_id
func (repo ColumnHandler) ColumnDetail(c echo.Context) error {
	columnId := c.Param("column_id")
	column, err := repo.ColumnRepo.GetColumn(columnId)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response{
			Code:    http.StatusNotFound,
			Message: "not found column_id",
		})
	}
	return c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "",
		Data:    column,
	})
}

// [GET] /column/:board_id
func (repo ColumnHandler) GetAllColumns(c echo.Context) error {
	boardId := c.Param("board_id")
	columns, err := repo.ColumnRepo.GetAllColumnsOfBoard(boardId)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response{
			Code:    http.StatusNotFound,
			Message: "not found board_id",
		})
	}
	return c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "",
		Data:    columns,
	})
}

// [PATCH] /column/update
func (repo ColumnHandler) UpdateColumn(c echo.Context) error {
	var column models.Column
	if err := c.Bind(&column); err != nil {
		return c.JSON(http.StatusNotFound, models.Response{
			Code:    http.StatusNotFound,
			Message: "can't mapping column",
		})
	}
	if err := repo.ColumnRepo.UpdateColumn(column); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: "failed to update",
		})
	}
	return c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "",
		Data:    column,
	})
}
