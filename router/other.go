package router

import (
	"github.com/gin-gonic/gin"
)

// HandleOther -- Route for handling other
func HandleOther(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "You need some authentication first!",
	})
}
