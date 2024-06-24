package seeder

import (
	"github.com/arioprima/cari_kampus_api/models"
	"github.com/arioprima/cari_kampus_api/pkg"
	"gorm.io/gorm"
	"log"
	"time"
)

func SeedUsers(db *gorm.DB) {
	//log.Println("Seeding users")
	users := []models.ModelAuth{
		{
			ID:        "019047ca-f542-7cf7-8471-453c5045417c",
			Nama:      "admin",
			Email:     "admin@gmail.com",
			Password:  pkg.HashPassword("admin1234"),
			RoleId:    "019047ca-f542-7182-8b6b-7978f905dfe7",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "019047ca-f542-795e-a5e0-486d50eeb38e",
			Nama:      "user",
			Email:     "user@gmail.com",
			Password:  pkg.HashPassword("user1234"),
			RoleId:    "019047ca-f542-71fe-9de6-c4919ed5c9ff",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	}

	//log.Println("users", users)

	for _, user := range users {
		result := db.Create(&user)
		if result.Error != nil {
			log.Fatalf("Failed to seed user data: %v", result.Error)
		}
		//log.Printf("User seeded successfully: %+v", user)
	}
}
