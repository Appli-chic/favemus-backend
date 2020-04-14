package main

import (
	"github.com/Favemus/config"
)

func main() {
	// Config the server
	config.LoadConfiguration()
	config.InitDB()

	r := InitRouter()
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
