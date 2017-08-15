package main

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func main() {

	// Use the default router from gin
	router = gin.Default()

	// Use the templates from the templates folder
	router.LoadHTMLGlob("templates/*")

	initializeRoutes()

	// start the application
	router.Run()
}
