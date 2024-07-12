package repository

import (
	"context"
	"trello-api/models"
)

type UserRepository interface {
	SaveUser(context context.Context, user models.User) (models.User, error)
	CheckUser(context context.Context, reqLogin models.ReqLogin) (models.User, error)
	GetUser(context context.Context, userId string) (models.User, error)
}
