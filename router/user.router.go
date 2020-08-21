package router

import (
	"fmt"
	"golang-restapi/model"
	"golang-restapi/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserData -- Retrieve user data
func GetUserData(c *gin.Context) {
	userID := uint64(c.MustGet("userID").(float64))
	fmt.Println("GetUserData userID", userID)
	var user model.User = repository.GetUserByID(userID)
	c.JSON(http.StatusOK, gin.H{
		"message": "Nice to see you bruh!",
		"data": gin.H{
			"_id":      user.ID,
			"username": user.Username,
			"password": user.Password,
		},
	})
}
