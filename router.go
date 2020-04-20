package main

import (
	"github.com/Favemus/controller"
	"github.com/Favemus/util"
	"github.com/gin-gonic/gin"
)

// Creates all the routes
func InitRouter() *gin.Engine {
	authController := controller.NewAuthController()
	userController := controller.NewUserController()

	r := gin.Default()

	// Adding routes
	api := r.Group("/api")
	{
		// Auth routes
		api.POST("/auth/signup", authController.SignUp)
		api.POST("/auth/login", authController.Login)
		api.POST("/auth/check/email", authController.CheckEmail)
		api.POST("/auth/refresh", authController.RefreshAccessToken)

		// Protected routes
		loggedInGroup := api.Group("/")
		loggedInGroup.Use(util.AuthenticationRequired())
		{
			loggedInGroup.GET("/users/:id", userController.GetUser)
		}
	}

	return r
}
