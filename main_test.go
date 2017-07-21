package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}

	// Start up application and open database connections
	a.Initialize(
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"),
		os.Getenv("TEST_DB_HOST"))

	// Check to see if the tableCreationQuery is able to successfully run
	ensureTableExists()

	// Run the web server
	code := m.Run()

	// Clear the database of any and all test data
	clearTable()

	// Exit the test
	os.Exit(code)
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/posts", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got %s", body)
	}

}

func TestGetNonExistentPost(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", "/post/13", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusNotFound, response.Code)

	var m map[string]string
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["error"] != "Post not found" {
		t.Errorf("Expected the 'error' key of the response to be set to 'Post not found'")
	}

}

func TestCreatePost(t *testing.T) {
	clearTable()

	payload := []byte(`{"name":"test post","author":"test author","body":"test content"}`)

	req, _ := http.NewRequest("POST", "/post", bytes.NewBuffer(payload))
	response := executeRequest(req)

	checkResponseCode(t, http.StatusCreated, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)
	if m["name"] != "test post" {
		t.Errorf("Expected post title to be 'test post'. Got %v", m["name"])
	}
	if m["author"] != "test author" {
		t.Errorf("Expected post author to be 'test author'. Got %v", m["author"])
	}
	if m["body"] != "test content" {
		t.Errorf("Expected post content to be 'test content'. Got %v", m["body"])
	}
	if m["id"] != 1.0 {
		t.Errorf("Expected post ID to be '1'. Got %v", m["id"])
	}
}

func TestGetPost(t *testing.T) {
	clearTable()
	addPost(1)

	req, _ := http.NewRequest("GET", "/post/1", nil)
	response := executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)
}

func TestUpdatePost(t *testing.T) {
	clearTable()
	addPost(1)

	req, _ := http.NewRequest("GET", "/post/1", nil)
	response := executeRequest(req)
	var originalPost map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &originalPost)

	payload := []byte(`{"name":"test post - updated name","author":"test author - updated name","body":"test content - updated content"}`)

	req, _ = http.NewRequest("PUT", "/post/1", bytes.NewBuffer(payload))
	response = executeRequest(req)

	checkResponseCode(t, http.StatusOK, response.Code)

	var m map[string]interface{}
	json.Unmarshal(response.Body.Bytes(), &m)

	if m["name"] == originalPost["name"] {
		t.Errorf("Expected post title to be %v. Got %v", originalPost["name"], m["name"])
	}
	if m["author"] == originalPost["author"] {
		t.Errorf("Expected author to be %v. Got %v", originalPost["author"], m["author"])
	}
	if m["body"] == originalPost["body"] {
		t.Errorf("Expected body to be %v. Got %v", originalPost["body"], m["body"])
	}
	if m["id"] != originalPost["id"] {
		t.Errorf("Expected the id to remain the same (%v). Got %v", originalPost["id"], m["id"])
	}
}

func TestDeletePost(t *testing.T) {
	clearTable()
	addPost(1)

	req, _ := http.NewRequest("GET", "/product/1", nil)
	response := executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("DELETE", "/product/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusOK, response.Code)

	req, _ = http.NewRequest("GET", "/product/1", nil)
	response = executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, response.Code)
}

func addPost(count int) {
	if count < 1 {
		count = 1
	}

	for i := 0; i < count; i++ {
		a.DB.Exec("INSERT INTO posts(name, author, body) VALUES($1, $2, $3)", "Post "+strconv.Itoa(i), "Author "+strconv.Itoa(i), "Test content for post "+strconv.Itoa(i))
	}
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM posts")
	a.DB.Exec("ALTER SEQUENCE posts_id_seq RESTART WITH 1")
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	a.Router.ServeHTTP(rr, req)

	return rr
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. Got %d\n", expected, actual)
	}
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS posts
(
id SERIAL,
name TEXT NOT NULL,
author TEXT NOT NULL,
body TEXT NOT NULL,
datecreated TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
dateupdated TIMESTAMPTZ NOT NULL DEFAULT current_timestamp,
CONSTRAINT posts_pkey PRIMARY KEY (id)
)`
