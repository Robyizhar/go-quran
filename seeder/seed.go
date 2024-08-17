package seeder

import (
	"go-quran/models"

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

	// Insert sample data
	products := []models.Product{
		{Name: "Product 1", Description: "Creating an Engine instance with the Logger and Recovery middleware already attached"},
		{Name: "Product 2", Description: "Creating an Engine instance with the Logger and Recovery middleware already attached"},
		{Name: "Product 3", Description: "Creating an Engine instance with the Logger and Recovery middleware already attached"},
	}

	for _, product := range products {
		db.Create(&product)
	}
}
