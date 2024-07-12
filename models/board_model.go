package models

import "time"

type Board struct {
	BoardId     string    `json:"boardId,omitempty" db:"board_id,omitempty"`
	Title       string    `json:"title,omitempty" db:"title,omitempty" validate:"required"`
	Description string    `json:"description,omitempty" db:"description,omitempty" validate:"required"`
	Type        string    `json:"type,omitempty" db:"type,omitempty" validate:"required"`
	CreatedAt   time.Time `json:"-" db:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"-" db:"updated_at,omitempty"`
}
