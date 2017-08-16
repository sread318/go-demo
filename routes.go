package main

func initializeRoutes() {
	router.GET("/", showIndexPage)
	router.GET("/articles/view/:article_id", getArticle)
}
