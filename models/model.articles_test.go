package models

import "testing"

func TestGetAllArticles(t *testing.T) {
	alist := GetAllArticles()

	// Check that alist is the same length as the
	// global article list count
	if len(alist) != len(ArticleList) {
		t.Fail()
	}

	// Check each article to make sure they match
	// the global article list
	for i, v := range alist {
		if v.Content != ArticleList[i].Content ||
			v.Title != ArticleList[i].Title ||
			v.ID != ArticleList[i].ID {
			t.Fail()
			break
		}
	}
}
