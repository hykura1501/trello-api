package repositoryImpl

import (
	"context"
	"log"
	"trello-api/database"
	"trello-api/models"
	"trello-api/repository"
)

type BoardRepositoryImpl struct {
	sql *database.SQL
}

func NewBoardRepository(sql *database.SQL) repository.BoardRepository {
	return &BoardRepositoryImpl{
		sql: sql,
	}
}

func (repo BoardRepositoryImpl) SaveBoard(context context.Context, board models.Board) error {
	statement := `
		INSERT INTO boards(board_id, title, description, type, created_at, updated_at)
		VALUES (:board_id, :title, :description, :type, :created_at, :updated_at)
	`
	if _, err := repo.sql.Db.NamedExecContext(context, statement, board); err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
func (repo BoardRepositoryImpl) GetBoard(boardId string) (models.Board, error) {
	var board models.Board
	statement := `
		SELECT * FROM boards WHERE board_id = ?
	`
	if err := repo.sql.Db.Get(&board, statement, boardId); err != nil {
		return board, err
	}
	return board, nil
}
