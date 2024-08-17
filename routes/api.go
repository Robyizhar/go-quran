package routes

import (
	"go-quran/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRoutes configures the routes for the application
func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	{
		productRoutes := api.Group("/products")
		{
			productRoutes.GET("/", controllers.Index)
			productRoutes.GET("/:id", controllers.Show)
			productRoutes.POST("/", controllers.Store)
			productRoutes.PUT("/:id", controllers.Update)
			productRoutes.DELETE("/:id", controllers.Delete)
		}
	}
}
