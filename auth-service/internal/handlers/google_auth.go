package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
    "auth-service/internal/services"
)

func GoogleLogin(c *gin.Context) {
    var body struct {
        IDToken string `json:"id_token"`
    }
    if err := c.ShouldBindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    payload, err := services.VerifyGoogleToken(
        body.IDToken,
        services.GoogleClientID,
    )
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }

    // find or create user (DB logic omitted for brevity)

    token := services.GenerateJWT(user)

    c.JSON(http.StatusOK, gin.H{
        "token": token,
        "user":  user,
    })
}
