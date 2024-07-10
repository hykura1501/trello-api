package models

import "time"

type User struct {
	UserId    string    `json:"userId,omitempty"`
	FullName  string    `json:"fullName,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"-"`
	Avatar    string    `json:"avatar,omitempty"`
	Birthday  time.Time `json:"birthday,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Address   string    `json:"address,omitempty"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
