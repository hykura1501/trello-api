package models

import "time"

type Column struct {
	BoardId     string    `json:"boardId,omitempty" db:"board_id,omitempty" validate:"required"`
	ColumnId    string    `json:"columnId,omitempty" db:"column_id,omitempty"`
	Title       string    `json:"title,omitempty" db:"title,omitempty" validate:"required"`
	ColumnOrder int       `json:"columnOrder,omitempty" db:"column_order,omitempty" `
	CreatedAt   time.Time `json:"-" db:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"-" db:"updated_at,omitempty"`
}
