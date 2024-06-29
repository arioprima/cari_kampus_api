package repositories

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/arioprima/cari_kampus_api/config"
	"github.com/arioprima/cari_kampus_api/models"
	"github.com/arioprima/cari_kampus_api/pkg"
	"github.com/arioprima/cari_kampus_api/schemas"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

type RepositoryLogin interface {
	Login(ctx context.Context, tx *gorm.DB, input *schemas.SchemaAuth) (*models.ModelAuth, *schemas.SchemaDatabaseError)
}

type repositoryLoginImpl struct {
	DB  *gorm.DB
	Log *logrus.Logger
}

func NewRepositoryLoginImpl(log *logrus.Logger, db *gorm.DB) RepositoryLogin {
	return &repositoryLoginImpl{
		DB:  db,
		Log: log,
	}
}

func (r *repositoryLoginImpl) Login(ctx context.Context, tx *gorm.DB, input *schemas.SchemaAuth) (*models.ModelAuth, *schemas.SchemaDatabaseError) {
	var (
		user models.ModelAuth
		err  error
	)

	// Begin transaction if tx is not nil
	if tx == nil {
		tx = r.DB.WithContext(ctx).Debug().Begin()
	}

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		} else if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// Check if user exists
	if err = tx.Where("email = ?", input.Email).Preload("UserRole").First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, &schemas.SchemaDatabaseError{
				Code: http.StatusNotFound,
				Type: "error_01",
			}
		}
		return nil, &schemas.SchemaDatabaseError{
			Code: http.StatusInternalServerError,
			Type: "error_02",
		}
	}

	// Compare passwords
	comparePassword := pkg.ComparePassword(user.Password, input.Password)
	if comparePassword != nil {
		return nil, &schemas.SchemaDatabaseError{
			Code: http.StatusUnauthorized,
			Type: "error_03",
		}
	}

	// Create user session
	configs, _ := config.LoadConfig(".")
	hashed := sha256.New()
	hashed.Write([]byte(configs.TokenSecret + time.Now().String()))
	token := hex.EncodeToString(hashed.Sum(nil))
	user.Token = token

	session := models.UserSession{
		UserID:    user.ID,
		Token:     token,
		LastLogin: time.Now(),
		ExpiredAt: pkg.CalculateExpiration(time.Now().Add(configs.TokenExpired).Unix()),
	}

	if err := tx.Create(&session).Error; err != nil {
		return nil, &schemas.SchemaDatabaseError{
			Code: http.StatusInternalServerError,
			Type: "error_02",
		}
	}

	log.Println("tokenRepository", token)

	return &user, nil
}
