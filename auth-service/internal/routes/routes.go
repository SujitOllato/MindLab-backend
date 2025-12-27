package routes

import (
    "github.com/gin-gonic/gin"
    "auth-service/internal/handlers"
)

func RegisterRoutes(r *gin.Engine) {
    auth := r.Group("/api/v1/auth")
    {
        auth.POST("/google", handlers.GoogleLogin)
        auth.POST("/signup", handlers.EmailSignup)
        auth.POST("/login", handlers.EmailLogin)
    }
}
