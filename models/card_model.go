package models

import (
	"time"
)

type Card struct {
	BoardId     string    `json:"boardId,omitempty" db:"board_id,omitempty" validate:"required"`
	ColumnId    string    `json:"columnId,omitempty" db:"column_id,omitempty" validate:"required"`
	CardId      string    `json:"cardId,omitempty" db:"card_id,omitempty"`
	Title       string    `json:"title,omitempty" db:"title,omitempty" validate:"required"`
	Description string    `json:"description,omitempty" db:"description,omitempty"`
	CardOrder   int       `json:"cardOrder,omitempty" db:"card_order,omitempty" `
	Thumbnail   string    `json:"thumbnail,omitempty" db:"thumbnail,omitempty" `
	CreatedAt   time.Time `json:"-" db:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"-" db:"updated_at,omitempty"`
}

type FileAttachment struct {
	FileUrl   string    `db:"file_url" json:"fileUrl"`
	BoardId   string    `db:"board_id" json:"boardId"`
	ColumnId  string    `db:"column_id" json:"columnId"`
	FileName  string    `db:"file_name" json:"fileName"`
	UserId    string    `db:"user_id" json:"userId"`
	CardId    string    `db:"card_id" json:"cardId"`
	CreatedAt time.Time `json:"-" db:"created_at,omitempty"`
	UpdatedAt time.Time `json:"-" db:"updated_at,omitempty"`
}
