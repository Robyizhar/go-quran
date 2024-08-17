package seeder

import (
	"go-quran/models"
	"strconv"

	"gorm.io/gorm"
)

// Seed function to insert data into the database
func Seed(db *gorm.DB) {
	// Check if products already exist
	var count int64
	db.Model(&models.Product{}).Count(&count)
	if count > 0 {
		return // Data already seeded
	}

	for i := 0; i < 10000; i++ {
		db.Create(&models.Product{
			Name:        "Product " + strconv.Itoa(i+1),
			Description: "Creating an Engine instance with the Logger and Recovery middleware already attached" + strconv.Itoa(i+1),
		})
	}
}
