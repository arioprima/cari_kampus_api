package services

import (
	"context"
	"github.com/arioprima/cari_kampus_api/config"
	"github.com/arioprima/cari_kampus_api/models"
	"github.com/arioprima/cari_kampus_api/pkg"
	repositories "github.com/arioprima/cari_kampus_api/repositories/auth"
	"github.com/arioprima/cari_kampus_api/schemas"
	"log"
	"time"
)

type ServiceLogin interface {
	LoginService(ctx context.Context, input *schemas.SchemaAuth) (*models.ModelAuth, *schemas.SchemaDatabaseError)
}

type serviceLoginImpl struct {
	repository repositories.RepositoryLogin
}

func NewServiceLoginImpl(repository repositories.RepositoryLogin) ServiceLogin {
	return &serviceLoginImpl{repository: repository}
}

func (s *serviceLoginImpl) LoginService(ctx context.Context, input *schemas.SchemaAuth) (*models.ModelAuth, *schemas.SchemaDatabaseError) {
	//TODO implement me
	var schema schemas.SchemaAuth
	schema.Email = input.Email
	schema.Password = input.Password

	res, err := s.repository.Login(ctx, &schema)
	if err != nil {
		return nil, err
	}

	configs, _ := config.LoadConfig(".")
	log.Println("Configs", configs)

	accessTokenData := map[string]interface{}{
		"id":      res.ID,
		"email":   res.Email,
		"nama":    res.Nama,
		"role_id": res.RoleId,
	}

	token, tokenErr := pkg.GenerateToken(accessTokenData, configs.TokenSecret, configs.TokenExpired)

	if tokenErr != nil {
		return nil, &schemas.SchemaDatabaseError{
			Code: 500,
			Type: "error_04",
		}
	}

	expired := time.Now().Add(configs.TokenExpired).Unix()
	log.Println("Expired", expired)
	log.Println("Expired", configs.TokenExpired)

	res.Auth.AccessToken = token
	res.Auth.Type = "Bearer"
	res.Auth.ExpiredAt = pkg.CalculateExpiration(expired)

	return res, nil
}
