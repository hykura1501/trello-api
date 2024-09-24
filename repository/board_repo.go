package repository

import (
	"context"
	"trello-api/models"
)

type BoardRepository interface {
	SaveBoard(context context.Context, board models.Board) error
	GetBoard(boardId string) (models.Board, error)
	InsertUser(boardId, userId, role string) error
}
