package controllers

import (
	"go-quran/database"
	"go-quran/models"
	"go-quran/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {

	var data []models.Product

	// Fetch all data from the database
	result := database.DB.Find(&data)
	if result.Error != nil {
		utils.ErrorResponse(c, "Failed to retrieve data", http.StatusInternalServerError)
		return
	}

	// Send response with the data data
	utils.SuccessResponse(c, data)

}

func Show(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam) // Convert the ID from string to int
	if err != nil {
		utils.ErrorResponse(c, "Invalid ID format", http.StatusBadRequest)
		return
	}

	var product models.Product
	result := database.DB.First(&product, id) // Fetch the product by ID
	if result.Error != nil {
		utils.ErrorResponse(c, "Product not found", http.StatusNotFound)
		return
	}

	utils.SuccessResponse(c, product)
}

func Store(c *gin.Context) {
	data := []string{"Product1", "Product2"}
	utils.SuccessResponse(c, data)
}

func Update(c *gin.Context) {
	data := []string{"Product1", "Product2"}
	utils.SuccessResponse(c, data)
}

func Delete(c *gin.Context) {
	data := []string{"Product1", "Product2"}
	utils.SuccessResponse(c, data)
}
