package main

import (
	"github.com/Favemus/config"
	"github.com/gin-gonic/gin"
)

func main() {
	// Config the server
	config.LoadConfiguration()
	config.InitDB()

	r := InitRouter()

	// Add midllewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
