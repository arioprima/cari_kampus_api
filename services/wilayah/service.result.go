package services

import (
	"github.com/arioprima/cari_kampus_api/models"
	repositories "github.com/arioprima/cari_kampus_api/repositories/wilayah"
	"github.com/arioprima/cari_kampus_api/schemas"
)

type ServiceResult interface {
	ResultWilayahService() (*[]models.Wilayah, schemas.SchemaDatabaseError)
}

type ServiceResultImpl struct {
	repository repositories.RepositoryResult
}

func NewServiceResultImpl(repository repositories.RepositoryResult) ServiceResult {
	return &ServiceResultImpl{repository: repository}

}

func (s *ServiceResultImpl) ResultWilayahService() (*[]models.Wilayah, schemas.SchemaDatabaseError) {
	//TODO implement me
	res, err := s.repository.GetWilayah()
	//log.Println("res", res)
	return res, err
}
