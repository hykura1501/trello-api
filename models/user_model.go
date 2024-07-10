package models

import "time"

type User struct {
	UserId    string    `json:"userId,omitempty" db:"user_id,omitempty"`
	FullName  string    `json:"fullName,omitempty" db:"full_name,omitempty"`
	Email     string    `json:"email,omitempty" db:"email,omitempty"`
	Password  string    `json:"-" db:"password,omitempty"`
	Avatar    string    `json:"avatar,omitempty" db:"avatar,omitempty"`
	Birthday  time.Time `json:"birthday,omitempty" db:"birthday,omitempty"`
	Phone     string    `json:"phone,omitempty" db:"phone,omitempty"`
	Address   string    `json:"address,omitempty" db:"address,omitempty"`
	CreatedAt time.Time `json:"-" db:"created_at,omitempty"`
	UpdatedAt time.Time `json:"-" db:"updated_at,omitempty"`
	Token     string    `json:"token,omitempty" db:"-"`
}

type ReqRegister struct {
	FullName string `json:"fullName" validate:"required"`
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type ReqLogin struct {
	Email    string `json:"email" validate:"required" db:"email"`
	Password string `json:"password" validate:"required" db:"password"`
}
