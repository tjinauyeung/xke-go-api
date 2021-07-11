package main

import "github.com/google/uuid"

var Users []User = []User{
	{ID: uuid.New(), FirstName: "John", LastName: "Doe", Email: "john@doe.com"},
	{ID: uuid.New(), FirstName: "Sally", LastName: "Sue", Email: "sally@sue.com"},
	{ID: uuid.New(), FirstName: "Mikey", LastName: "Mike", Email: "mikey@mike.com"},
}
