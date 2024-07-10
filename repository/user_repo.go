package repository

import (
	"context"
	"trello-api/models"
)

type UserRepository interface {
	SaveUser(context context.Context, user models.User) (models.User, error)
}
