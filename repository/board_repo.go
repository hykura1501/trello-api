package repository

import (
	"context"
	"trello-api/models"
)

type BoardRepository interface {
	SaveBoard(context context.Context, board models.Board) error
}
