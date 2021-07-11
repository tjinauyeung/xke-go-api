package main

import (
	"gorm.io/gorm"
)

func CreateUser(u *User, db *gorm.DB) ([]User, error) {
	tx := db.Create(u)
	if tx.Error != nil {
		return nil, &Error{msg: "Could not create user", code: ErrorConflict}
	}

	uu, err := GetUsers(db)
	if err != nil {
		return nil, err
	}

	return uu, nil
}

func GetUser(id string, db *gorm.DB) (*User, error) {
	var user *User

	tx := db.First(&user, id)
	if tx.Error != nil {
		return nil, &Error{msg: "Could not find user", code: ErrorNotFound}
	}

	return user, nil
}

func GetUsers(db *gorm.DB) ([]User, error) {
	var users []User

	tx := db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}
