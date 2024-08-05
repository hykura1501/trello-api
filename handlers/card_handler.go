package handlers

import (
	"log"
	"net/http"
	"time"
	"trello-api/models"
	"trello-api/repository"
	"trello-api/utils"

	validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type CardHandler struct {
	CardRepo repository.CardRepository
}

// [POST] /card/new
func (repo CardHandler) NewCard(c echo.Context) error {
	var card models.Card
	if err := c.Bind(&card); err != nil {
		log.Println(err.Error())
		return c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	validate := validator.New()
	if err := validate.Struct(card); err != nil {
		return c.JSON(http.StatusNotAcceptable, models.Response{
			Code:    http.StatusNotAcceptable,
			Message: err.Error(),
		})
	}
	// insert to db
	card.CardId = uuid.New().String()
	card.CreatedAt = time.Now()
	card.UpdatedAt = time.Now()
	if err := repo.CardRepo.SaveCard(&card); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "card created",
		Data:    card,
	})
}

// [GET] /card/detail/:card_id
func (repo CardHandler) CardDetail(c echo.Context) error {
	cardId := c.Param("card_id")
	card, err := repo.CardRepo.GetCard(cardId)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response{
			Code:    http.StatusNotFound,
			Message: "not found card_id",
		})
	}
	return c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "",
		Data:    card,
	})
}

// [GET] /card/:column_id
func (repo CardHandler) GetAllCards(c echo.Context) error {
	columnId := c.Param("column_id")
	cards, err := repo.CardRepo.GetAllCardsOfColumn(columnId)
	if err != nil {
		return c.JSON(http.StatusNotFound, models.Response{
			Code:    http.StatusNotFound,
			Message: "not found column_id",
		})
	}
	return c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "",
		Data:    cards,
	})
}

// [PATCH] /card/update
func (repo CardHandler) UpdateCard(c echo.Context) error {
	var card models.Card
	if err := c.Bind(&card); err != nil {
		return c.JSON(http.StatusNotFound, models.Response{
			Code:    http.StatusNotFound,
			Message: "can't mapping card",
		})
	}
	if err := repo.CardRepo.UpdateCard(card); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Code:    http.StatusBadRequest,
			Message: "failed to update",
		})
	}
	return c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "",
		Data:    card,
	})
}

// [POST] /card/attachment/new
func (repo CardHandler) NewAttachment(c echo.Context) error {
	formHeader, err := c.FormFile("file")
	boardId := c.FormValue("boardId")
	columnId := c.FormValue("columnId")
	cardId := c.FormValue("cardId")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	formFile, err := formHeader.Open()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	uploadUrl, err := utils.ImageUploadHelper(formFile)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	//handle after uploadFile
	reqAttachment := models.FileAttachment{
		BoardId:   boardId,
		ColumnId:  columnId,
		CardId:    cardId,
		UserId:    "723b8ef5-09a8-4265-bfb9-810c1d83611f",
		FileName:  formHeader.Filename,
		FileUrl:   uploadUrl,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if err := repo.CardRepo.SaveAttachment(reqAttachment); err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    reqAttachment,
	})
}

// [POST] /card/attachment
func (repo CardHandler) GetAllAttachments(c echo.Context) error {
	var card models.Card
	if err := c.Bind(&card); err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	attachments, err := repo.CardRepo.GetAllAttachments(card)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "success",
		Data:    attachments,
	})
}

// [DELETE] /card/attachment/delete
func (repo CardHandler) DeleteAttachment(c echo.Context) error {
	var attachment models.FileAttachment
	if err := c.Bind(&attachment); err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	if err := repo.CardRepo.DeleteAttachment(attachment); err != nil {
		return c.JSON(http.StatusInternalServerError, models.Response{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, models.Response{
		Code:    http.StatusOK,
		Message: "success",
	})
}
