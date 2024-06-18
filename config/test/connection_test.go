package test

import (
	"log"
	"testing"

	"github.com/arioprima/cari_kampus_api/config"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var connectionDb *gorm.DB
var err error

func init() {
	cfg, err := config.LoadConfig("../../")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}
	connectionDb, err = config.OpenConnection(&cfg)
	if err != nil {
		log.Fatalf("Error opening connection: %v", err)
	}
}

func TestConnectDB(t *testing.T) {
	assert.Nil(t, err)
	assert.NotNil(t, connectionDb)
}
