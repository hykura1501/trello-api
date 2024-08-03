package repository

import "trello-api/models"

type CardRepository interface {
	SaveCard(card *models.Card) error
	GetCard(cardId string) (models.Card, error)
	GetAllCardsOfColumn(columnId string) ([]models.Card, error)
	UpdateCard(cards models.Card) error
	SaveAttachment(attachment models.FileAttachment) error
}
