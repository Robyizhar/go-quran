package main

import (
	"go-quran/database"
	"go-quran/routes"
	"go-quran/seeder"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.InitDatabase()

	seeder.Seed(database.DB)

	routes.SetupRoutes(r)
	r.Run()
}
