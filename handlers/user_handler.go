package handlers

import (
	"log"
	"net/http"
	"trello-api/models"
	"trello-api/repository"
	"trello-api/security"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserRepo repository.UserRepository
}

func (repo UserHandler) Register(c echo.Context) error {
	var reqRegister models.ReqRegister
	if err := c.Bind(&reqRegister); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	user := models.User{
		UserId:   uuid.New().String(),
		FullName: reqRegister.FullName,
		Email:    reqRegister.Email,
		Password: security.HashAndSalt([]byte(reqRegister.Password)), // Hash Password
	}
	//Update to DB
	user, err := repo.UserRepo.SaveUser(c.Request().Context(), user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	//Return Token
	token, err := security.GenerateToken(user)
	if err != nil {
		return c.JSON(http.StatusRequestTimeout, models.Response{
			Code:    http.StatusRequestTimeout,
			Message: err.Error(),
		})
	}
	user.Token = token
	return c.JSON(http.StatusOK, models.Response{
		Code: http.StatusOK,
		Data: user,
	})
}

func (repo UserHandler) Login(c echo.Context) error {
	return c.JSON(http.StatusOK, "Login")
}
