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

	// Routes for User
	userRoute := r.Group("/user")
	userRoute.Use(middleware.AuthMiddleware())
	{
		userRoute.GET("/", router.GetUserData)
	}

	// Routes for Service
	serviceRoute := r.Group("/service")
	serviceRoute.Use(middleware.AuthMiddleware())
	{
		serviceRoute.POST("/", router.CreateServiceRequest)
	}

	r.Run()
}
