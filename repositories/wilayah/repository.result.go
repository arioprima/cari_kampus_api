package wilayah

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
	var wilayah []models.Wilayah
	err := r.DB.Find(&wilayah).Error
	if err != nil {
		return nil, schemas.SchemaDatabaseError{
			Code: http.StatusInternalServerError,
			Type: "error_01",
		}
	}
	return &wilayah, schemas.SchemaDatabaseError{}
}
