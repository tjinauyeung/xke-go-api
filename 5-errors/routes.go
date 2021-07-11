package main

import (
	"github.com/gorilla/mux"
)

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/users", CreateUserHandler).Methods("POST")
	r.HandleFunc("/users", GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", GetUserHandler).Methods("GET")
}
