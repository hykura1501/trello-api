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

func (repo CardRepositoryImpl) SaveCard(card *models.Card) error {
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

func (repo CardRepositoryImpl) UpdateCard(cards models.Card) error {
	//update title
	statement :=
		`	
		UPDATE cards SET 
			title= CASE WHEN length(:title)=0 THEN title ELSE :title END,
			description= CASE WHEN length(:description)=0 THEN description ELSE :description END
		WHERE board_id=:board_id and column_id=:column_id and card_id=:card_id
	`
	if _, err := repo.sql.Db.NamedExec(statement, cards); err != nil {
		log.Println(err)
		return err
	}
	return nil
}
func (repo CardRepositoryImpl) SaveAttachment(attachment models.FileAttachment) error {
	statement := "INSERT INTO card_attachments(board_id, column_id, card_id, user_id, url_file, created_at, updated_at) VALUES (:board_id, :column_id, :card_id, :user_id, :url_file, :created_at, :updated_at)"
	if _, err := repo.sql.Db.NamedExec(statement, attachment); err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (repo CardRepositoryImpl) GetAllAttachments(card models.Card) ([]models.FileAttachment, error) {
	statement := `
		SELECT * FROM card_attachments WHERE card_id=? AND column_id=? and board_id=? ORDER BY created_at DESC
	`
	var attachments []models.FileAttachment
	if err := repo.sql.Db.Select(&attachments, statement, card.CardId, card.ColumnId, card.BoardId); err != nil {
		return attachments, err
	}
	return attachments, nil
}
