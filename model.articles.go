package main

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

var articleList = []article{
	article{Title: "Article 1", Content: "Content for article one"},
	article{Title: "Article 2", Content: "Content for article two"},
}

func getAllArticles() []article {
	return articleList
}
