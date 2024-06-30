package repositories

import (
	"github.com/arioprima/cari_kampus_api/models"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

func FinByToken(userId string, db *gorm.DB) (string, error) {
	var userToken []models.UserSession
	err := db.Debug().Order("created_at desc").
		Select("token").Where("user_id = ? and expired_at >= NOW()", userId).
		First(&userToken).Error

	if err != nil {
		logrus.Error(err)
		return "", err
	}

	return userToken[0].Token, nil
}
