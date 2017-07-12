package main

import (
	"database/sql"
	"errors"
	"time"
)

type post struct {
	ID          int       `json:"id"`
	Title       string    `json:"name"`
	Author      string    `json:"author"`
	Body        string    `json:"body"`
	DatePosted  time.Time `json:"dateposted"`
	DateUpdated time.Time `json:"dateupdated"`
}

func (p *post) getPost(db *sql.DB) error {
	return errors.New("Not yet implemented")
}

func (p *post) updatePost(db *sql.DB) error {
	return errors.New("Not yet implemented")
}

func (p *post) createPost(db *sql.DB) error {
	return errors.New("Not yet implemented")
}

func (p *post) deletePost(db *sql.DB) error {
	return errors.New("Not yet implemented")
}

func (p *post) getPosts(db *sql.DB, start, count int) ([]post, error) {
	return nil, errors.New("Not yet implemented")
}
