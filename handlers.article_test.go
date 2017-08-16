package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestShowIndexPageUnauthenticated(t *testing.T) {
	r := getRouter()

	r.GET("/", showIndexPage)

	req, _ := http.NewRequest("GET", "/", nil)

	testHTTPResponse(t, r, req, func(w *httptest.ResponseRecorder) bool {

		// Check to see if the status code is 200
		statusOK := w.Code == http.StatusOK

		// Read the page into memory
		p, err := ioutil.ReadAll(w.Body)

		// Make sure there isn't an error and look for the proper title
		pageOK := err == nil && strings.Index(string(p), "<title>Home Page</title>") > 0

		return statusOK && pageOK
	})
}
