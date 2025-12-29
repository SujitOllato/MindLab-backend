package handlers

import (
	"context"
	"net/http"

	"auth-service/internal/services"

	"github.com/gin-gonic/gin"
	"google.golang.org/api/idtoken"
)

type GoogleLoginRequest struct {
	IDToken string `json:"id_token"`
}

func GoogleLogin(c *gin.Context) {
	var req GoogleLoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	payload, err := idtoken.Validate(
		context.Background(),
		req.IDToken,
		services.GetGoogleClientID(),
	)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Google token"})
		return
	}

	user, err := services.FindOrCreateGoogleUser(
		payload.Claims["email"].(string),
		payload.Claims["name"].(string),
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	token, _ := services.GenerateJWT(user.ID, user.Email)

	c.JSON(http.StatusOK, gin.H{
		"token": token,
		"user":  user,
	})
}
