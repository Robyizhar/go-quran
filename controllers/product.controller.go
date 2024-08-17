package controllers

import (
	"go-quran/database"
	"go-quran/models"
	"go-quran/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListProduct(c *gin.Context) {

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	offset := (page - 1) * limit

	var data []models.Product

	var total int64
	database.DB.Model(&models.Product{}).Count(&total)

	result := database.DB.Limit(limit).Offset(offset).Find(&data)
	if result.Error != nil {
		utils.ErrorResponse(c)
		return
	}

	meta := gin.H{"page": page, "limit": limit, "total": total}
	utils.SuccessResponse(c, data, http.StatusOK, "Berhasil menampilkan data !", meta)

}

func ShowProduct(c *gin.Context) {
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

func StoreProduct(c *gin.Context) {
	data := []string{"Product1", "Product2"}
	utils.SuccessResponse(c, data)
}

func UpdateProduct(c *gin.Context) {
	data := []string{"Product1", "Product2"}
	utils.SuccessResponse(c, data)
}

func DeleteProduct(c *gin.Context) {
	data := []string{"Product1", "Product2"}
	utils.SuccessResponse(c, data)
}
