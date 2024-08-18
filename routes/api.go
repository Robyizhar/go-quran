package routes

import (
	"go-quran/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {

	externalRoutes := r.Group("/external")
	{
		kenemagRoutes := externalRoutes.Group("/kemenag")
		{
			kenemagRoutes.GET("/syncronize", controllers.SyncronizeQuran)
		}
	}

	api := r.Group("/api")
	{
		product := api.Group("/products")
		{
			product.GET("/", controllers.ListProduct)
			product.GET("/:id", controllers.ShowProduct)
			product.POST("/", controllers.StoreProduct)
			product.PUT("/:id", controllers.UpdateProduct)
			product.DELETE("/:id", controllers.DeleteProduct)
		}

		surah := api.Group("/surah")
		{
			surah.GET("/", controllers.ListSurah)
			surah.GET("/ayah", controllers.ListAyah)
		}
	}
}
