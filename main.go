package main

import (
	"golang-restapi/middleware"
	"golang-restapi/repository"
	"golang-restapi/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize database connection
	repository.InitDb()

	// Routes for authentication
	authRoute := r.Group("/auth")
	{
		authRoute.POST("/register", router.Register)
		authRoute.POST("/login", router.Login)
	}

	// Routes for other
	otherRoute := r.Group("/other")
	otherRoute.Use(middleware.AuthMiddleware())
	{
		otherRoute.GET("/", router.HandleOther)
	}

	r.Run()
}
