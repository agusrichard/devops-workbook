package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
	token := c.Request.Header.Get("Authorization")
	if len(token) == 0 {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"message": "You shall not pass!",
		})
	} else {
		token = strings.Split(token, "Bearer ")[1]
		getPayload(token)
		c.Next()
	}
}

func getPayload(tokenString string) (interface{}, error) {
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	fmt.Println(token)
	if err != nil {
		panic(err)
		// return nil, err
	}

	// do something with decoded claims
	fmt.Println("claims", claims)
	for key, val := range claims {
		fmt.Printf("Key: %v, value: %v\n", key, val)
	}
	return nil, nil
}
