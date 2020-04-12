package main

import (
	"github.com/Favemus/controller"
	"github.com/gin-gonic/gin"
)

// Creates all the routes
func InitRouter() *gin.Engine {
	r := gin.Default()

	userController := controller.NewUserController()

	api := r.Group("/api")
	{
		api.GET("/users/:id", userController.GetUser)
	}

	return r
}
