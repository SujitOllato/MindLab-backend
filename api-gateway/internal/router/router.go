package router

import (
	"github.com/gin-gonic/gin"
	"api-gateway/internal/middleware"
	"api-gateway/internal/proxy"
)

func SetupRoutes(r *gin.Engine) {

	r.Use(middleware.Logger())
	r.Use(middleware.RateLimit())

	api := r.Group("/api")
	{
		api.POST("/auth/*path", proxy.AuthProxy)

		secured := api.Group("/")
		secured.Use(middleware.JWTAuth())

		secured.Any("/assessment/*path", proxy.AssessmentProxy)
		secured.Any("/scoring/*path", proxy.ScoringProxy)
		secured.Any("/report/*path", proxy.ReportProxy)
		secured.Any("/recommendation/*path", proxy.RecommendationProxy)
		secured.Any("/payment/*path", proxy.PaymentProxy)
	}
}
