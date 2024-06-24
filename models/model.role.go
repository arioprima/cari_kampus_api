package models

import "time"

type Role struct {
	ID        string    `json:"id" validate:"uuid"`
	Nama      string    `json:"nama" validate:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (role *Role) TableName() string {
	return "user_role"
}
