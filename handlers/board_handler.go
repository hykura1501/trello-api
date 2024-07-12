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

type BoardHandler struct {
	BoardRepo repository.BoardRepository
}

func (repo BoardHandler) NewBoard(c echo.Context) error {
	var board models.Board
	if err := c.Bind(&board); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	validate := validator.New()
	if err := validate.Struct(board); err != nil {
		return c.JSON(http.StatusNotAcceptable, models.Response{
			Code:    http.StatusNotAcceptable,
			Message: err.Error(),
		})
	}
	// insert to db
	board.BoardId = uuid.New().String()
	board.CreatedAt = time.Now()
	board.UpdatedAt = time.Now()
	if err := repo.BoardRepo.SaveBoard(c.Request().Context(), board); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "board created",
		Data:    board,
	})
}
