package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get all users")
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	fmt.Fprintf(w, "Get user: %s", params["id"])
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Create a user")
}
