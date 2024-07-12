package handlers

import (
	"log"
	"net/http"
	"time"
	"trello-api/models"
	"trello-api/repository"
	"trello-api/security"

	validator "github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserRepo repository.UserRepository
}

// [POST] /user/register
func (repo UserHandler) Register(c echo.Context) error {
	var reqRegister models.ReqRegister
	if err := c.Bind(&reqRegister); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	validate := validator.New()
	if err := validate.Struct(reqRegister); err != nil {
		return c.JSON(http.StatusNotAcceptable, models.Response{
			Code:    http.StatusNotAcceptable,
			Message: err.Error(),
		})
	}
	user := models.User{
		UserId:    uuid.New().String(),
		FullName:  reqRegister.FullName,
		Email:     reqRegister.Email,
		Password:  security.HashAndSalt([]byte(reqRegister.Password)), // Hash Password
		Birthday:  time.Now(),
		UpdatedAt: time.Now(),
		CreatedAt: time.Now(),
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
	token, expiresAt, err := security.GenerateToken(user)
	if err != nil {
		return c.JSON(http.StatusRequestTimeout, models.Response{
			Code:    http.StatusRequestTimeout,
			Message: err.Error(),
		})
	}
	security.WriteCookie(c, "token", token, expiresAt)

	return c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Data:    user,
		Message: "successful account registration",
	})
}

// [POST] /user/login
func (repo UserHandler) Login(c echo.Context) error {
	var reqLogin models.ReqLogin
	if err := c.Bind(&reqLogin); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	validate := validator.New()
	if err := validate.Struct(reqLogin); err != nil {
		return c.JSON(http.StatusNotAcceptable, models.Response{
			Code:    http.StatusNotAcceptable,
			Message: err.Error(),
		})
	}

	user, err := repo.UserRepo.CheckUser(c.Request().Context(), reqLogin)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response{
			Code:    http.StatusNotFound,
			Message: err.Error(),
		})
	}

	isTheSamePass := security.ComparePasswords(user.Password, []byte(reqLogin.Password))
	if !isTheSamePass {
		return c.JSON(http.StatusNotFound, models.Response{
			Code:    http.StatusNotFound,
			Message: "Incorrect Password",
		})
	}

	//Token
	token, expiresAt, err := security.GenerateToken(user)

	if err != nil {
		return c.JSON(http.StatusRequestTimeout, models.Response{
			Code:    http.StatusRequestTimeout,
			Message: err.Error(),
		})
	}

	security.WriteCookie(c, "token", token, expiresAt)

	return c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Data:    user,
		Message: "logged in successfully",
	})
}

// [GET] /user/profile
func (repo UserHandler) Profile(c echo.Context) error {
	tokenData := c.Get("user").(*jwt.Token)
	claims := tokenData.Claims.(*models.JWTCustomsClaims)
	userId := claims.UserId
	user, err := repo.UserRepo.GetUser(c.Request().Context(), userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response{
			Code:    http.StatusNotFound,
			Message: "Not Found",
		})
	}
	return c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Data:    user,
		Message: "request successfully",
	})
}
