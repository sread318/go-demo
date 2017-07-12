package main

import (
	"log"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}

	a.Initialize(
		os.Getenv("TEST_DB_USERNAME"),
		os.Getenv("TEST_DB_PASSWORD"),
		os.Getenv("TEST_DB_NAME"),
		os.Getenv("TEST_DB_HOST"))

	ensureTableExists()

	code := m.Run()

	clearTable()

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
