package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tylerconlee/go-demo/models"
)

func ShowIndexPage(c *gin.Context) {
	// Load all articles for the index page
	articles := models.GetAllArticles()

	// using the HTML method for the context, load the index page
	render(c, gin.H{
		"title":   "Homepage",
		"payload": articles,
	}, "index.html")
}

func GetArticle(c *gin.Context) {
	if articleID, err := strconv.Atoi(c.Param("article_id")); err == nil {
		if article, err := models.GetArticleByID(articleID); err == nil {
			render(c, gin.H{
				"title":   article.Title,
				"payload": article,
			}, "article.html")

		} else {
			c.AbortWithError(http.StatusNotFound, err)
		}
	} else {
		c.AbortWithStatus(http.StatusNotFound)
	}
}
