package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jesseinvent/go-otp-service/api"
)

func main() {
	router := gin.Default()

	// Initialize config
	app := api.Config{Router: router}

	// routes
	app.Routes()

	router.Run(":8000")
}
