package controllers

import (
	"encoding/json"
	"fmt"
	"go-quran/database"
	"go-quran/models"
	"go-quran/services"
	"go-quran/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SyncronizeQuran(c *gin.Context) {
	httpService := services.NewHttpService()

	url := "https://web-api.qurankemenag.net"
	headers := map[string]string{}

	surahURL := fmt.Sprintf("%s/quran-surah", url)
	responseSurahReq, err := httpService.Get(surahURL, headers)
	if err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to fetch data from external API")
		return
	}

	var responseSurah map[string]interface{}
	if err := json.Unmarshal(responseSurahReq, &responseSurah); err != nil {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Failed to parse API responseSurah")
		return
	}

	data, ok := responseSurah["data"].([]interface{})
	if !ok {
		utils.ErrorResponse(c, http.StatusInternalServerError, "Invalid data format")
		return
	}

	for _, item := range data {
		itemMap, ok := item.(map[string]interface{})
		if !ok {
			continue // Skip if the item is not in the expected format
		}

		surah := models.Surah{
			Id:             int16(itemMap["id"].(float64)), // Use float64 for numbers from JSON
			IndonesianName: itemMap["latin"].(string),
			ArabicName:     itemMap["arabic"].(string),
			Meaning:        itemMap["translation"].(string),
			Location:       itemMap["location"].(string),
			TotalAyah:      int(itemMap["num_ayah"].(float64)),
		}

		// Insert into the database
		result := database.DB.Save(&surah)
		ayahURL := fmt.Sprintf("%s/quran-ayah?surah=%d", url, surah.Id)
		responseAyahReq, err := httpService.Get(ayahURL, headers)
		log.Println(ayahURL)
		if err == nil {
			var responseAyah map[string]interface{}
			if err := json.Unmarshal(responseAyahReq, &responseAyah); err == nil {
				dataAyah, ok := responseAyah["data"].([]interface{})
				if ok {
					for _, itemAyah := range dataAyah {
						itemAyahMap, ok := itemAyah.(map[string]interface{})
						if !ok {
							log.Printf("Not Inserted Ayah ")
							continue // Skip if the item is not in the expected format
						}
						ayah := models.Ayah{
							Id:             int16(itemAyahMap["id"].(float64)), // Use float64 for numbers from JSON
							SurahId:        uint(surah.Id),
							NumberAyah:     getInt(itemAyahMap["ayah"]),
							Juz:            getInt(itemAyahMap["juz"]),
							Arabic:         getString(itemAyahMap["arabic"]),
							Kitabah:        getString(itemAyahMap["kitabah"]),
							Latin:          getString(itemAyahMap["latin"]),
							Translation:    getString(itemAyahMap["translation"]),
							Interpretation: getString(itemAyahMap["footnotes"]),
						}
						database.DB.Save(&ayah)
					}
				} else {
					log.Printf("Not Inserted Ayah ")
				}
			} else {
				log.Printf("Not Inserted Ayah ")
			}
		} else {
			log.Printf("Not Inserted Ayah %s", err)
		}
		if result.Error != nil {
			log.Printf("Failed to insert data for surah %s: %v", surah.IndonesianName, result.Error)
		} else {
			log.Printf("Successfully inserted surah %s", surah.IndonesianName)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   responseSurah["data"],
	})
}

func getInt(value interface{}) int {
	if value == nil {
		return 0
	}
	if v, ok := value.(float64); ok {
		return int(v)
	}
	return 0
}

func getString(value interface{}) string {
	if value == nil {
		return ""
	}
	return value.(string)
}
