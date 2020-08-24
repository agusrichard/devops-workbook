package handler

import (
	"golang-restapi/model"
	"golang-restapi/repository"
	"golang-restapi/utils"

	"github.com/gin-gonic/gin"
)

// Register router function
func Register(c *gin.Context) {
	var user model.User
	c.BindJSON(&user)
	var err error
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		utils.ResponseBadRequest(c, "Invalid username")
	}
	ok := repository.CreateUser(&user)
	if ok {
		utils.ResponseSuccess(c, "Register success", gin.H{
			"username": user.Username,
			"password": user.Password,
		})
	} else {
		utils.ResponseBadRequest(c, "Username has been used")
	}
}

// Login router function
func Login(c *gin.Context) {
	var loginUser model.User
	var user model.User
	var err error
	var token string
	c.BindJSON(&loginUser)
	user, err = repository.GetUserByUsername(loginUser.Username)
	if err != nil {
		panic(err)
	}
	if utils.CheckPasswordHash(loginUser.Password, user.Password) {
		token, err = utils.CreateToken(user.ID)
		if err != nil {
			panic(err)
		}
		utils.ResponseSuccess(c, "Login success", gin.H{
			"user": gin.H{
				"_id":      user.ID,
				"username": user.Username,
				"password": user.Password,
			},
			"token": token,
		})
	} else {
		utils.ResponseBadRequest(c, "Invalid username or password")
	}
}
