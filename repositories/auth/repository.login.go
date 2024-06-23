package repositories

import (
	"context"
	"errors"
	"github.com/arioprima/cari_kampus_api/models"
	"github.com/arioprima/cari_kampus_api/pkg"
	"github.com/arioprima/cari_kampus_api/schemas"
	"gorm.io/gorm"
	"net/http"
)

type RepositoryLogin interface {
	Login(ctx context.Context, input *schemas.SchemaAuth) (*models.ModelAuth, *schemas.SchemaDatabaseError)
}

type repositoryLoginImpl struct {
	DB *gorm.DB
}

func NewRepositoryLoginImpl(db *gorm.DB) RepositoryLogin {
	return &repositoryLoginImpl{DB: db}
}

func (r *repositoryLoginImpl) Login(ctx context.Context, input *schemas.SchemaAuth) (*models.ModelAuth, *schemas.SchemaDatabaseError) {
	//TODO implement me
	var user models.ModelAuth
	db := r.DB.WithContext(ctx).Debug().Model(&user)

	user.Email = input.Email
	user.Password = input.Password

	if err := db.Where("email = ?", input.Email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &schemas.SchemaDatabaseError{
				Code:    http.StatusNotFound,
				Type:    "error_01",
				Error:   err.Error(),
				Message: "User not found",
			}
		}
		return nil, &schemas.SchemaDatabaseError{
			Code:    http.StatusInternalServerError,
			Type:    "error_02",
			Error:   err.Error(),
			Message: "Internal server error",
		}
	}

	comparePassword := pkg.ComparePassword(user.Password, input.Password)
	if comparePassword != nil {
		return nil, &schemas.SchemaDatabaseError{
			Code:    http.StatusUnauthorized,
			Type:    "error_03",
			Error:   comparePassword.Error(),
			Message: "Password not match",
		}
	}

	return &user, nil
}
