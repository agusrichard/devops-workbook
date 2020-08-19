package router

import (
	"golang-restapi/models"
	"golang-restapi/repository"
	"golang-restapi/utils"

	"github.com/gin-gonic/gin"
)

// Register router function
func Register(c *gin.Context) {
	if c.Request.Method == "POST" {
		var user models.User
		c.BindJSON(&user)
		var err error
		user.Password, err = utils.HashPassword(user.Password)
		if err != nil {
			panic(err)
		}
		repository.CreateUser(&user)
		c.JSON(200, gin.H{
			"message": "Register Success",
			"data": gin.H{
				"_id":      user.ID,
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
		if utils.CheckPasswordHash(loginUser.Password, user.Password) {
			token, err := utils.CreateToken(user.ID)
			if err != nil {
				panic(err)
			}
			c.JSON(200, gin.H{
				"message": "Login Success",
				"data": gin.H{
					"user": gin.H{
						"username": user.Username,
						"password": user.Password,
					},
					"token": token,
				},
			})
		} else {
			c.JSON(400, gin.H{
				"message": "Failed to login",
			})
		}
	}
}
