package main

import (
	"fmt"
	"log"
	"os"

	"github.com/arioprima/cari_kampus_api/config"
	"github.com/arioprima/cari_kampus_api/db/seeder"
)

func main() {
	configPath := "."
	// Load configuration
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	// Open database connection
	db, err := config.OpenConnection(&cfg)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	log.Println("Database connection opened")
	// Seed data
	seeder.SeedRole(db)
	seeder.SeedUsers(db)

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error getting DB object: %v", err)
	}
	defer sqlDB.Close()

	// Selesai
	fmt.Println("Successfully inserted data")

	// Exit program
	os.Exit(0)
}
