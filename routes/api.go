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
		productRoutes := api.Group("/products")
		{
			productRoutes.GET("/", controllers.ListProduct)
			productRoutes.GET("/:id", controllers.ShowProduct)
			productRoutes.POST("/", controllers.StoreProduct)
			productRoutes.PUT("/:id", controllers.UpdateProduct)
			productRoutes.DELETE("/:id", controllers.DeleteProduct)
		}
	}
}
