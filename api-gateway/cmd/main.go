package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"api-gateway/internal/router"
)

func main() {
	_ = godotenv.Load()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := gin.New()
	r.Use(gin.Recovery())

	router.SetupRoutes(r)

	log.Println("🚀 API Gateway running on port", port)
	r.Run(":" + port)
}
