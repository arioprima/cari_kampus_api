package test

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViper(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("yaml")
	config.AddConfigPath("../../")
	config.AutomaticEnv()

	//read config
	err := config.ReadInConfig()
	assert.Nil(t, err)

	assert.Equal(t, "cari_kampus_api", config.GetString("app.name"))
	assert.Equal(t, "mysql", config.GetString("database.type"))
}
