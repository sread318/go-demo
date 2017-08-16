package main

import (
	"net/http"
	"strconv"

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

func getArticle(c *gin.Context) {
	if articleID, err := strconv.Atoi(c.Param("article_id")); err != nil {
		if article, err := getArticleByID(articleID); err == nil {
			c.HTML(
				http.StatusOK,
				"article.html",
				gin.H{
					"title":   article.Title,
					"payload": article.Content,
				},
			)
		}
	}

}
