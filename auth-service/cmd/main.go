package main

import (
	"log"

	"auth-service/internal/config"
	"auth-service/internal/database"
	"auth-service/internal/routes"
	"auth-service/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	cfg := config.Load()

	if err := database.Connect(cfg.DBUrl); err != nil {
		log.Fatal("Database connection failed:", err)
	}

	services.Init(cfg)

	r := gin.Default()
	routes.RegisterRoutes(r)

	r.Run(":8081")
}
