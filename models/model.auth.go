package models

import "time"

type ModelAuth struct {
	ID string `json:"id" gorm:"primary_key;column:id;<-:create"`
	Password string `json:"password,omitempty" gorm:"column:password;type:varchar(255)"`
	Nama string `json:"nama" gorm:"column:nama"`
	Email string `json:"email" gorm:"column:email"`
	RoleId string `json:"role_id" gorm:"column:role_id"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt 	time.Time `json:"updated_at" gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func (auth *ModelAuth) TableName() string {
	return "users"
}