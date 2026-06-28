package main

import (
	"log"

	"ego/database"
	"ego/handlers"
	"ego/templates"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi Database
	database.Init()

	// Setup Gin
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	// Setup Templates & Static
	templates.Setup(r)

	// Setup Routes
	handlers.SetupRoutes(r)

	// Jalankan di port 8080
	log.Println("🚀 ShadowSelf berjalan di :8080")
	r.Run(":8080")
}
