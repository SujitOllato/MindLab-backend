package handlers

import "github.com/gin-gonic/gin"

func EmailSignup(c *gin.Context) {
	c.JSON(501, gin.H{"message": "Email signup coming next"})
}

func EmailLogin(c *gin.Context) {
	c.JSON(501, gin.H{"message": "Email login coming next"})
}
