package router

import (
	"github.com/gin-gonic/gin"
)

// GetUserData -- Retrieve user data
func GetUserData(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "GetUserData route",
	})
}
