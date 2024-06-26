package wilayah

import (
	"github.com/arioprima/cari_kampus_api/models"
	"github.com/arioprima/cari_kampus_api/schemas"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
	"net/http"
	"time"
)

var c = cache.New(5*time.Minute, 10*time.Minute)

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

	// Check cache
	if x, found := c.Get("wilayah"); found {
		wilayah = x.([]models.Wilayah)
		return &wilayah, schemas.SchemaDatabaseError{}
	}

	// Fetch from database if not in cache
	err := r.DB.Find(&wilayah).Error
	if err != nil {
		return nil, schemas.SchemaDatabaseError{
			Code: http.StatusInternalServerError,
			Type: "error_01",
		}
	}

	// Set cache
	c.Set("wilayah", wilayah, cache.DefaultExpiration)
	return &wilayah, schemas.SchemaDatabaseError{}
}
