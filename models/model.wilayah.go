package models

type Wilayah struct {
	Kode string `json:"kode" gorm:"primary_key;column:kode"`
	Nama string `json:"nama" gorm:"column:nama"`
}

func (w *Wilayah) TableName() string {
	return "wilayah"
}
