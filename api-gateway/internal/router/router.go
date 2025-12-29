package router

import (
	"os"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"api-gateway/internal/middleware"
	"api-gateway/internal/proxy"
)

func SetupRoutes(r *gin.Engine) {

	r.Use(middleware.Logger())
	r.Use(middleware.RateLimit())

	api := r.Group("/api")
	{
		// Gateway health
		api.GET("/health", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status":  "ok",
				"message": "API Gateway is running",
			})
		})

		// Health check for Auth Service
		api.GET("/health/auth", func(c *gin.Context) {
			client := http.Client{Timeout: 3 * time.Second}
			resp, err := client.Get(os.Getenv("AUTH_SERVICE") + "/health")
			if err != nil {
				log.Println("Auth Service health check failed:", err)
				c.JSON(http.StatusServiceUnavailable, gin.H{
					"service": "auth-service",
					"status":  "unreachable",
				})
				return
			}
			defer resp.Body.Close()

			body, _ := ioutil.ReadAll(resp.Body)
			c.Data(resp.StatusCode, "application/json", body)
		})

		// Public routes (reverse proxy)
		api.POST("/auth/*path", proxy.AuthProxy)

		// Secured routes
		secured := api.Group("/")
		secured.Use(middleware.JWTAuth())

		secured.Any("/assessment/*path", proxy.AssessmentProxy)
		secured.Any("/scoring/*path", proxy.ScoringProxy)
		secured.Any("/report/*path", proxy.ReportProxy)
		secured.Any("/recommendation/*path", proxy.RecommendationProxy)
		secured.Any("/payment/*path", proxy.PaymentProxy)
	}
}
