package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"api-gateway/internal/router"
)

func main() {
	r := gin.Default()

	// Setup all routes
	router.SetupRoutes(r)

	// Get port from env or default
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("API Gateway running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
