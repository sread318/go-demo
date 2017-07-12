package main

import (
	"log"
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

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM posts")
	a.DB.Exec("ALTER SEQUENCE posts_id_seq RESTART WITH 1")
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
