package main

import (
	"errors"
)

type post struct {
	ID    		int `json:"id"`
	Title 		string `json:"name"`
	Author 		string `json:"author"`
	Body  		string `json:"body"`
	DatePosted 	int64 `json:"dateposted"`
	DateUpdated int64 `json:"dateupdated"`
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

func (p *post) getPosts(db *sql.DB) error {
	return errors.New("Not yet implemented")
}