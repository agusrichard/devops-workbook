package main

import (
	"golang-restapi/config"
	"golang-restapi/handler"
	"golang-restapi/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize database connection
	config.InitDb()

	// Routes for authentication
	authRoute := r.Group("/auth")
	{
		authRoute.POST("/register", handler.Register)
		authRoute.POST("/login", handler.Login)
	}

	// Routes for User
	userRoute := r.Group("/user")
	userRoute.Use(middleware.AuthMiddleware())
	{
		userRoute.GET("/", handler.GetUserData)
	}

	// Routes for Service
	serviceRoute := r.Group("/service")
	serviceRoute.Use(middleware.AuthMiddleware())
	{
		serviceRoute.POST("/", handler.CreateServiceRequest)
		serviceRoute.GET("/", handler.GetServices)
	}

	r.Run()
}
