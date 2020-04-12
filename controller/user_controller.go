package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
}

func NewUserController() *UserController {
	return &UserController{}
}

func (u *UserController) GetUser(c *gin.Context) {
	// Send the user information
	c.JSON(http.StatusOK, gin.H{
		"id":   0,
		"name": "test",
	})
}
