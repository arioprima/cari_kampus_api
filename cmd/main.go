package main

import (
	"github.com/arioprima/cari_kampus_api/config"
	"github.com/arioprima/cari_kampus_api/routes"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	loadConfig, err := config.LoadConfig(".")
	if err != nil {
		log.Println("Error loading config:", err)
		return
	}

	db, err := config.OpenConnection(&loadConfig)
	if err != nil {
		log.Println("Error opening database connection:", err)
		return
	}

	router := routes.SetupAuthRoutes(db)
	port := loadConfig.PORT
	err = router.Run(":" + port)
	if err != nil {
		log.Println("Could not run server:", err)
	}
}
