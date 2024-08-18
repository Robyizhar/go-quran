package main

import (
	"go-quran/database"
	"go-quran/routes"
	"go-quran/seeder"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.InitDatabase()

	seeder.Seed(database.DB)

	routes.SetupRoutes(r)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}

	r.Run(":" + port)
}
