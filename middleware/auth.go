package middleware

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware -- Authentication Middleware
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		validateToken(c)
		c.Next()
	}
}

func validateToken(c *gin.Context) {
	reqToken := c.Request.Header.Get("Authorization")
	fmt.Printf("reqToken %v", reqToken)
	if len(reqToken) == 0 {
		c.AbortWithStatusJSON(300, gin.H{
			"message": "You shall not pass!",
		})
	}
	splitToken := strings.Split(reqToken, "Bearer ")
	fmt.Printf("splitToken %v", splitToken)
	reqToken = splitToken[1]
	fmt.Printf("reqTokenv %v", reqToken)

	c.Next()
}
