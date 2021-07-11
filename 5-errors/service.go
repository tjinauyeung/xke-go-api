package main

import (
	"github.com/google/uuid"
)

func CreateUser(u *User) ([]User, error) {
	// generate ID before appending
	Users = append(Users, User{
		ID:        uuid.New(),
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Email:     u.Email,
	})
	return Users, nil
}

func GetUser(uid uuid.UUID) (*User, error) {
	for _, u := range Users {
		if u.ID == uid {
			return &u, nil
		}
	}

	return nil, &Error{msg: "Could not find user", code: ErrorNotFound}
}

func GetUsers() []User {
	return Users
}
