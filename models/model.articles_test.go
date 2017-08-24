package models

import "testing"

func TestGetAllArticles(t *testing.T) {
	alist := getAllArticles()

	// Check that alist is the same length as the
	// global article list count
	if len(alist) != len(articleList) {
		t.Fail()
	}

	// Check each article to make sure they match
	// the global article list
	for i, v := range alist {
		if v.Content != articleList[i].Content ||
			v.Title != articleList[i].Title ||
			v.ID != articleList[i].ID {
			t.Fail()
			break
		}
	}
}
