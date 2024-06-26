package repositories

import (
	"github.com/arioprima/cari_kampus_api/models"
	"github.com/arioprima/cari_kampus_api/schemas"
	"gorm.io/gorm"
	"net/http"
)

type RepositoryResult interface {
	GetWilayah() (*[]models.Wilayah, schemas.SchemaDatabaseError)
}

type repositoryWilayahImpl struct {
	DB *gorm.DB
}

func NewRepositoryWilayahImpl(db *gorm.DB) RepositoryResult {
	return &repositoryWilayahImpl{DB: db}
}

func (r *repositoryWilayahImpl) GetWilayah() (*[]models.Wilayah, schemas.SchemaDatabaseError) {
	//TODO implement me
	var wilayah []models.Wilayah
	db := r.DB.Model(&wilayah)

	errorCode := make(chan schemas.SchemaDatabaseError, 1)

	resultStudent := db.Debug().Find(&wilayah)

	if resultStudent.RowsAffected < 1 {
		errorCode <- schemas.SchemaDatabaseError{
			Code: http.StatusNotFound,
			Type: "error_01",
		}
		return &wilayah, <-errorCode
	}
	return &wilayah, <-errorCode
}
