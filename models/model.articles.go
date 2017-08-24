package models

import "errors"

type Article struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// TODO: convert to loadArticle to articleList and saveArticle (append to articleList)
var ArticleList = []Article{
	Article{ID: 1, Title: "Article 1", Content: "Content for article one"},
	Article{ID: 2, Title: "Article 2", Content: "Content for article two"},
}

func GetAllArticles() []Article {
	return ArticleList
}

func GetArticleByID(id int) (*Article, error) {
	for _, a := range ArticleList {
		if a.ID == id {
			return &a, nil
		}
	}
	return nil, errors.New("Article not found")
}
