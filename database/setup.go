package database

import (
	"go-quran/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := "host=127.0.0.1 user=postgres password=asdw1234 dbname=go-quran port=5432 sslmode=disable TimeZone=Asia/Singapore"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	log.Println("Connected to the database successfully!")

	database.AutoMigrate(&models.Users{}, &models.Product{}, &models.Surah{}, &models.Ayah{})
	DB = database
}
