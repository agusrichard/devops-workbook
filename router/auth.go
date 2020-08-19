package router

import (
	"golang-restapi/models"
	"golang-restapi/repository"

	"github.com/gin-gonic/gin"
)

// Register router function
func Register(c *gin.Context) {
	if c.Request.Method == "POST" {
		var user models.User
		c.BindJSON(&user)
		repository.CreateUser(&user)
		c.JSON(200, gin.H{
			"message": "Register Success",
			"data": gin.H{
				"username": user.Username,
				"password": user.Password,
			},
		})
	}

}

// Login router function
func Login(c *gin.Context) {
	if c.Request.Method == "POST" {
		var loginUser models.User
		c.BindJSON(&loginUser)
		var user models.User = repository.GetUserByUsername(loginUser.Username)
		c.JSON(200, gin.H{
			"message": "Login Success",
			"data": gin.H{
				"username": user.Username,
				"password": user.Password,
			},
		})
	}
}
