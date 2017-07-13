package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
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
