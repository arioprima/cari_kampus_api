package models

type Wilayah struct {
	Kode string `json:"id" gorm:"primary_key;column:kode"`
	Nama string `json:"nama" gorm:"column:kode"`
}

func (w *Wilayah) TableName() string {
	return "wilayah"
}