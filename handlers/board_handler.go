package handlers

import (
	"log"
	"net/http"
	"time"
	"trello-api/models"
	"trello-api/repository"

	validator "github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type BoardHandler struct {
	BoardRepo repository.BoardRepository
	UserRepo  repository.UserRepository
}

// [GET] /board/new
func (repo BoardHandler) NewBoard(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*models.JWTCustomsClaims)
	userId := claims.UserId

	//check user is existed in db
	_, err := repo.UserRepo.GetUser(c.Request().Context(), userId)
	if err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusUnauthorized, models.Response{
			Code:    http.StatusUnauthorized,
			Message: err.Error(),
		})
	}

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
	// insert board to db
	board.BoardId = uuid.New().String()
	board.CreatedAt = time.Now()
	board.UpdatedAt = time.Now()
	if err := repo.BoardRepo.SaveBoard(c.Request().Context(), board); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	// insert member to db
	if err := repo.BoardRepo.InsertUser(board.BoardId, userId, "admin"); err != nil {
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

// [GET] /board/detail/:board_id
func (repo BoardHandler) BoardDetail(c echo.Context) error {
	boardId := c.Param("board_id")
	board, err := repo.BoardRepo.GetBoard(boardId)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response{
			Code:    http.StatusNotFound,
			Message: "not found board_id",
		})
	}
	return c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "",
		Data:    board,
	})
}
