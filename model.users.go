package main

import "errors"

type user struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

// TODO: Add in loadUser/saveUser

func registerUser(username, password string) (*user, error) {
	return nil, errors.New("Placeholder")
}

func checkUsernameAvail(username string) bool {
	return false
}
