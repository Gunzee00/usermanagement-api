package config

import (
	"fmt"
	"log"
	"os"

	"github.com/Gunzee00/usermanagement_api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gagal koneksi database: %v", err)
	}

	// Auto migrate langsung di sini
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Gagal migrasi: %v", err)
	}

	DB = db
}
