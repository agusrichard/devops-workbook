package main

import (
	"golang-restapi/repository"
	"golang-restapi/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize database connection
	repository.InitDb()

	// Router for authentication
	r.POST("/auth/register", router.Register)
	r.GET("/auth/login", router.Login)

	r.Run()
}
