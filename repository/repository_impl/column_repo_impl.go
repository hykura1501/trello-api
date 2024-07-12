package repositoryImpl

import (
	"trello-api/database"
	"trello-api/repository"
)

type ColumnRepositoryImpl struct {
	sql *database.SQL
}

func NewColumnRepository(sql *database.SQL) repository.ColumnRepository {
	return &ColumnRepositoryImpl{
		sql: sql,
	}
}
