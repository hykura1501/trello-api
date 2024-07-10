package repositoryImpl

import (
	"trello-api/database"
	"trello-api/repository"
)

type UserRepositoryImpl struct {
	sql *database.SQL
}

func NewUserRepository(sql *database.SQL) repository.UserRepository {
	return &UserRepositoryImpl{
		sql: sql,
	}
}
