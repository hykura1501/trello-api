package repositoryImpl

import (
	"log"
	"trello-api/database"
	"trello-api/models"
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

func (repo ColumnRepositoryImpl) SaveColumn(column models.Column) error {
	maxOrder, err := database.GetMaxOrder(repo.sql, "columns", "board_id", column.BoardId)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	column.ColumnOrder = maxOrder + 1
	statement := `
	INSERT INTO columns(board_id, column_id, column_order, title, created_at, updated_at)
	VALUES (:board_id, :column_id, :column_order, :title, :created_at, :updated_at)
`
	if _, err := repo.sql.Db.NamedExec(statement, column); err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (repo ColumnRepositoryImpl) GetColumn(columnId string) (models.Column, error) {
	var column models.Column
	statement := `
		SELECT * FROM columns WHERE column_id = ?
	`
	if err := repo.sql.Db.Get(&column, statement, columnId); err != nil {
		return column, err
	}
	return column, nil
}

func (repo ColumnRepositoryImpl) GetAllColumnsOfBoard(boardId string) ([]models.Column, error) {
	var columns []models.Column
	statement := `
		SELECT * FROM columns WHERE board_id = ? ORDER BY column_order asc
	`
	if err := repo.sql.Db.Select(&columns, statement, boardId); err != nil {
		return columns, err
	}
	return columns, nil
}
