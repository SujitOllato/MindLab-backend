package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestEmailSignup(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/auth/signup", EmailSignup)

	body := map[string]string{
		"email":    "test@example.com",
		"password": "123456",
		"name":     "Test User",
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/auth/signup", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK && w.Code != 501 {
		t.Errorf("Expected status 200 or 501, got %d", w.Code)
	}
}

func TestEmailLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/auth/login", EmailLogin)

	body := map[string]string{
		"email":    "test@example.com",
		"password": "123456",
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusOK && w.Code != 501 {
		t.Errorf("Expected status 200 or 501, got %d", w.Code)
	}
}

func TestGoogleLogin_InvalidToken(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	router.POST("/auth/google", GoogleLogin)

	body := map[string]string{
		"id_token": "invalid_token",
	}
	jsonBody, _ := json.Marshal(body)

	req, _ := http.NewRequest("POST", "/auth/google", bytes.NewBuffer(jsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusUnauthorized && w.Code != http.StatusBadRequest {
		t.Errorf("Expected status 401 or 400 for invalid token, got %d", w.Code)
	}
}
