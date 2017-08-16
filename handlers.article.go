package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	// Load all articles for the index page
	articles := getAllArticles()

	// using the HTML method for the context, load the index page
	c.HTML(
		// set the status to 200 (OK)
		http.StatusOK,
		// load the index.html template
		"index.html",
		// pass the page specific information
		gin.H{
			"title":   "Home Page",
			"payload": articles,
		},
	)
}
