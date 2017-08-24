package models

import "errors"

type article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// TODO: convert to loadArticle to articleList and saveArticle (append to articleList)
var ArticleList = []article{
	article{ID: 1, Title: "Article 1", Content: "Content for article one"},
	article{ID: 2, Title: "Article 2", Content: "Content for article two"},
}

func GetAllArticles() []article {
	return ArticleList
}

func GetArticleByID(id int) (*article, error) {
	for _, a := range ArticleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
