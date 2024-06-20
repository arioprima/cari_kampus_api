package repositories

import (
	"context"
	"github.com/arioprima/cari_kampus_api/models"
	"github.com/arioprima/cari_kampus_api/schemas"
	"gorm.io/gorm"
)

type RepositoryLogin interface {
	Login(ctx context.Context, tx *gorm.DB, input *schemas.SchemaAuth) (*models.ModelAuth, *schemas.SchemaDatabaseError)
}

type repositoryLoginImpl struct {
	DB *gorm.DB
}

func NewRepositoryLoginImpl(db *gorm.DB) RepositoryLogin {
	return &repositoryLoginImpl{DB: db}
}

func (r *repositoryLoginImpl) Login(ctx context.Context, tx *gorm.DB, input *schemas.SchemaAuth) (*models.ModelAuth, *schemas.SchemaDatabaseError) {
	//TODO implement me
	panic("implement me")
}
