package repositories

import (
	"errors"
	"github.com/arioprima/cari_kampus_api/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func GetRoleNameById(roleId string, db *gorm.DB) (string, error) {
	var role []models.Role
	err := db.Debug().Select("nama").Where("id = ?", roleId).First(&role).Error
	if err != nil {
		logrus.Error(errors.New("role not found"))
		return "", errors.New("role not found")
	}
	return role[0].Nama, nil
}
