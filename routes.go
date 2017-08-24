package main

import "github.com/tylerconlee/go-demo/handlers"

func initializeRoutes() {
	router.GET("/", handlers.ShowIndexPage)
	router.GET("/article/view/:article_id", handlers.GetArticle)
}
