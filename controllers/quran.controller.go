package controllers

import (
	"fmt"
	"go-quran/database"
	"go-quran/models"
	"go-quran/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListSurah(c *gin.Context) {
	page, limit, orderBy, orderType, offset := utils.GetPaginationParameters(c)

	validOrderBys := map[string]bool{"id": true, "indonesian_name": true, "arabic_name": true, "meaning": true, "location": true, "total_ayah": true}
	orderBy = utils.ValidateOrderBy(orderBy, validOrderBys)

	var data []models.Surah
	var total int64

	database.DB.Model(&models.Surah{}).Count(&total)
	orderClause := fmt.Sprintf("%s %s", orderBy, orderType)

	result := database.DB.Order(orderClause).Limit(limit).Offset(offset).Find(&data)
	if result.Error != nil {
		utils.ErrorResponse(c)
		return
	}

	meta := gin.H{"page": page, "limit": limit, "total": total}
	utils.SuccessResponse(c, data, http.StatusOK, "Berhasil menampilkan data !", meta)
}

func ListAyah(c *gin.Context) {
	page, limit, orderBy, orderType, offset := utils.GetPaginationParameters(c)
	validOrderBys := map[string]bool{"id": true, "surah_id": true, "number_ayah": true, "juz": true, "arabic": true, "latin": true, "translation": true}
	orderBy = utils.ValidateOrderBy(orderBy, validOrderBys)

	surahID := c.Query("surah_id")
	if surahID == "" {
		log.Println(http.StatusBadRequest)
		utils.ErrorResponse(c, http.StatusBadRequest, "surah_id is required")
		return
	}

	var data []models.Ayah
	var total int64

	database.DB.Model(&models.Ayah{}).Where("surah_id = ?", surahID).Count(&total)
	orderClause := fmt.Sprintf("%s %s", orderBy, orderType)

	result := database.DB.Where("surah_id = ?", surahID).Order(orderClause).Limit(limit).Offset(offset).Find(&data)
	if result.Error != nil {
		utils.ErrorResponse(c)
		return
	}

	meta := gin.H{"page": page, "limit": limit, "total": total}
	utils.SuccessResponse(c, data, http.StatusOK, "Berhasil menampilkan data !", meta)
}
