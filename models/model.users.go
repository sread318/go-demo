package models

import "errors"

type User struct {
	Username string `json:"username"`
	Password string `json:"-"`
}

var UserList = []User{
	User{Username: "user1", Password: "pass1"},
	User{Username: "user2", Password: "pass2"},
	User{Username: "user3", Password: "pass3"},
}

// TODO: Add in loadUser/saveUser

func RegisterUser(username, password string) (*User, error) {
	return nil, errors.New("Placeholder")
}

func checkUsernameAvail(username string) bool {
	return false
}
