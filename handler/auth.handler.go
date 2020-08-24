package handler

import (
	"fmt"
	"golang-restapi/model"
	"golang-restapi/repository"
	"golang-restapi/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Register handler function
func Register(c *gin.Context) {
	var user model.User
	var err error
	c.BindJSON(&user)
	fmt.Println("Register user", len(user.Email), len(user.Password))
	if len(user.Email) == 0 || len(user.Password) == 0 {
		utils.ResponseBadRequest(c, "Please provide email and password")
		return
	}
	user.Password, err = utils.HashPassword(user.Password)
	if err != nil {
		utils.ResponseServerError(c)
		return
	}
	user.UUID = uuid.New().String()
	ok := repository.CreateUser(&user)
	if ok {
		utils.ResponseSuccess(c, "Register success", gin.H{
			"email":    user.Email,
			"password": user.Password,
			"uuid":     user.UUID,
		})
	} else {
		utils.ResponseBadRequest(c, "Email has been used")
	}
}

// Login handler function
func Login(c *gin.Context) {
	var loginUser model.User
	var user model.User
	var err error
	var token string
	c.BindJSON(&loginUser)
	if len(loginUser.Email) == 0 || len(loginUser.Password) == 0 {
		utils.ResponseBadRequest(c, "Please provide email and password")
		return
	}
	user, err = repository.GetUserByEmail(loginUser.Email)
	if err != nil {
		utils.ResponseServerError(c)
		return
	}
	if utils.CheckPasswordHash(loginUser.Password, user.Password) {
		token, err = utils.CreateToken(user.ID)
		if err != nil {
			utils.ResponseServerError(c)
			return
		}
		utils.ResponseSuccess(c, "Login success", gin.H{
			"user": gin.H{
				"_id":      user.ID,
				"username": user.Email,
				"password": user.Password,
			},
			"token": token,
		})
		return
	}
	utils.ResponseBadRequest(c, "Invalid username or password")
	return
}

// ConfirmAccount ...
func ConfirmAccount(c *gin.Context) {
	var data model.ConfirmData
	var user model.User
	var err error
	var ok bool
	c.BindJSON(&data)
	fmt.Println("data", data)
	if len(data.Email) == 0 || len(data.Password) == 0 || len(data.UUID) == 0 {
		utils.ResponseBadRequest(c, "Please provide email, password, and uuid")
	}
	user, err = repository.GetUserByEmail(data.Email)
	if err != nil {
		utils.ResponseServerError(c)
		return
	}
	fmt.Println("user", user)
	fmt.Println(user.UUID, data.UUID)
	fmt.Println(user.UUID == data.UUID)
	if user.UUID != data.UUID {
		utils.ResponseBadRequest(c, "Wrong confirmation code")
		return
	}
	ok, err = repository.ConfirmAccount(user.ID)
	if !ok || err != nil {
		utils.ResponseServerError(c)
		return
	}
	utils.ResponseSuccess(c, "Success to confirm account", gin.H{
		"email": user.Email,
	})
}
