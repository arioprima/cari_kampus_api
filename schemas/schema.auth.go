package schemas

import "time"

type SchemaAuth struct {
	ID        string    `json:"id" validate:"uuid"`
	Password  string    `json:"password" validate:"required,min=3,max=100"`
	Nama      string    `json:"nama" validate:"required"`
	Email     string    `json:"email" validate:"required,email"`
	RoleId    string    `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LoginResponse struct {
	ID        string    `json:"id"`
	Nama      string    `json:"nama"`
	Email     string    `json:"email"`
	RoleId    string    `json:"role_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
