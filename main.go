package main

import "github.com/gin-gonic/gin"

func main() {
	r := InitRouter()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	err := r.Run()

	if err != nil {
		panic(err)
	}
}
