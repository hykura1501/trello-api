package repository

import "trello-api/models"

type ColumnRepository interface {
	SaveColumn(column models.Column) error
	GetColumn(columnId string) (models.Column, error)
	GetAllColumnsOfBoard(boardId string) ([]models.Column, error)
}
