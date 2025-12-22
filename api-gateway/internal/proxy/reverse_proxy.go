package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/gin-gonic/gin"
)

func proxyRequest(target string, c *gin.Context) {
	remote, _ := url.Parse(target)
	proxy := httputil.NewSingleHostReverseProxy(remote)

	c.Request.URL.Path = c.Param("path")
	proxy.ServeHTTP(c.Writer, c.Request)
}

func AuthProxy(c *gin.Context) {
	proxyRequest(os.Getenv("AUTH_SERVICE"), c)
}

func AssessmentProxy(c *gin.Context) {
	proxyRequest(os.Getenv("ASSESSMENT_SERVICE"), c)
}

func ScoringProxy(c *gin.Context) {
	proxyRequest(os.Getenv("SCORING_SERVICE"), c)
}

func ReportProxy(c *gin.Context) {
	proxyRequest(os.Getenv("REPORT_SERVICE"), c)
}

func RecommendationProxy(c *gin.Context) {
	proxyRequest(os.Getenv("RECOMMENDATION_SERVICE"), c)
}

func PaymentProxy(c *gin.Context) {
	proxyRequest(os.Getenv("PAYMENT_SERVICE"), c)
}
