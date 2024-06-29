package services

import (
	"context"
	"github.com/arioprima/cari_kampus_api/config"
	"github.com/arioprima/cari_kampus_api/models"
	"github.com/arioprima/cari_kampus_api/pkg"
	repositories "github.com/arioprima/cari_kampus_api/repositories/auth"
	"github.com/arioprima/cari_kampus_api/schemas"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
	"time"
)

type ServiceLogin interface {
	LoginService(ctx context.Context, tx *gorm.DB, input *schemas.SchemaAuth) (*models.ModelAuth, *schemas.SchemaDatabaseError)
}

type serviceLoginImpl struct {
	repository repositories.RepositoryLogin
	Log        *logrus.Logger
}

func NewServiceLoginImpl(repository repositories.RepositoryLogin, logger *logrus.Logger) ServiceLogin {
	return &serviceLoginImpl{
		repository: repository,
		Log:        logger,
	}
}

func (s *serviceLoginImpl) LoginService(ctx context.Context, tx *gorm.DB, input *schemas.SchemaAuth) (*models.ModelAuth, *schemas.SchemaDatabaseError) {
	//TODO implement me
	var schema schemas.SchemaAuth
	schema.Email = input.Email
	schema.Password = input.Password

	res, err := s.repository.Login(ctx, tx, &schema)
	if err != nil {
		s.Log.Error(err)
		return nil, err
	}

	configs, _ := config.LoadConfig(".")
	log.Println("Configs", configs)

	log.Println("test", res.Token)

	accessTokenData := map[string]interface{}{
		"id":      res.ID,
		"email":   res.Email,
		"nama":    res.Nama,
		"token":   res.Token,
		"role_id": res.RoleId,
	}

	token, tokenErr := pkg.GenerateToken(accessTokenData, configs.TokenSecret, configs.TokenExpired)

	if tokenErr != nil {
		return nil, &schemas.SchemaDatabaseError{
			Code: 500,
			Type: "error_04",
		}
	}

	res.Auth.AccessToken = token
	res.Auth.Type = "Bearer"
	res.Auth.ExpiredAt = pkg.CalculateExpiration(time.Now().Add(configs.TokenExpired).Unix())

	return res, nil
}
