package main

import (
	"github.com/arioprima/cari_kampus_api/config"
	"github.com/arioprima/cari_kampus_api/routes"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	db, err := config.OpenConnection(&loadConfig)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	router := gin.Default()

	routes.SetupAuthRoutes(router, db)
	routes.SetupWilayahRoutes(router, db)

	port := loadConfig.PORT
	err = router.Run(":" + port)
	if err != nil {
		log.Fatalf("Could not run server: %v", err)
	}
}
