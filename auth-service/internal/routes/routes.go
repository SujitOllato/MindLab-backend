package routes

import (
	"auth-service/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {

	// Health check (for browser, load balancer, UAT checks)
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "ok",
			"service": "auth-service",
		})
	})

	auth := r.Group("/auth")
	{
		auth.POST("/google", handlers.GoogleLogin)
		auth.POST("/signup", handlers.EmailSignup)
		auth.POST("/login", handlers.EmailLogin)
	}
}
