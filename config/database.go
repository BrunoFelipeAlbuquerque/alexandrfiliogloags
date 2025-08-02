package config

import (
	"fmt"
	"log"
	"os"

	"imperial-fleet/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	dsn := os.Getenv("DB_DSN") // example: "host=localhost user=postgres password=postgres dbname=imperial port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to DB: ", err)
	}

	err = db.AutoMigrate(&models.Spaceship{}, &models.Armament{})
	if err != nil {
		log.Fatal("Failed to migrate DB: ", err)
	}

	DB = db
	fmt.Println("🚀 Connected to database.")
}
