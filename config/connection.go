package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func OpenConnection(config *Config) (*gorm.DB, error) {
	dialect := mysql.Open(config.Database.DBUser + ":" + config.Database.DBPass + "@tcp(" + config.Database.DBHost + ":" + config.Database.DBPort + ")/" + config.Database.DBName + "?charset=utf8&parseTime=True&loc=Asia%2FJakarta");
	db, err := gorm.Open(dialect, &gorm.Config{})
	if err != nil {
	return nil, err
	}
	return db, nil
}




