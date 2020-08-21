package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ResponseSuccess ...
func ResponseSuccess(c *gin.Context, message string, data gin.H) {
	c.JSON(http.StatusOK, gin.H{
		"message": message,
		"data":    data,
	})
}

// ResponseUnauthorized ...
func ResponseUnauthorized(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
		"message": "Unauthorized",
	})
}

// ResponseBadRequest ...
func ResponseBadRequest(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
		"message": message,
	})
}
