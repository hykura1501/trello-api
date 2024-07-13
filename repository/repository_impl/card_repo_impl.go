package repositoryImpl

import (
	"log"
	"trello-api/database"
	"trello-api/models"
	"trello-api/repository"
)

type CardRepositoryImpl struct {
	sql *database.SQL
}

func NewCardRepository(sql *database.SQL) repository.CardRepository {
	return &CardRepositoryImpl{
		sql: sql,
	}
}

func (repo CardRepositoryImpl) SaveCard(card models.Card) error {
	maxOrder, err := database.GetMaxOrder(repo.sql, "cards", "column_id", card.ColumnId)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	card.CardOrder = maxOrder + 1
	statement := `
	INSERT INTO cards(board_id, column_id, card_id, card_order, title, description, thumbnail, created_at, updated_at)
	VALUES (:board_id, :column_id, :card_id, :card_order,:title, :description, :thumbnail, :created_at, :updated_at)
`
	if _, err := repo.sql.Db.NamedExec(statement, card); err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (repo CardRepositoryImpl) GetCard(cardId string) (models.Card, error) {
	var card models.Card
	statement := `
		SELECT * FROM cards WHERE card_id = ?
	`
	if err := repo.sql.Db.Get(&card, statement, cardId); err != nil {
		return card, err
	}
	return card, nil
}

func (repo CardRepositoryImpl) GetAllCardsOfColumn(columnId string) ([]models.Card, error) {
	var card []models.Card
	statement := `
		SELECT * FROM cards WHERE column_id = ? ORDER BY card_order asc
	`
	if err := repo.sql.Db.Select(&card, statement, columnId); err != nil {
		return card, err
	}
	return card, nil
}
