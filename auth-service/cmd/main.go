package main

import (
    "github.com/gin-gonic/gin"
    "auth-service/internal/config"
    "auth-service/internal/database"
    "auth-service/internal/routes"
)

func main() {
    cfg := config.Load()
    db, _ := database.Connect(cfg.DBUrl)

    r := gin.Default()
    routes.RegisterRoutes(r)

    r.Run(":3000")
}
