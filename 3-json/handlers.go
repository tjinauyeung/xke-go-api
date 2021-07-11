package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
	response(w, Users, http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	uid := params["id"]

	for _, u := range Users {
		if u.ID == uid {
			json.NewEncoder(w).Encode(u)
			return
		}
	}

	response(w, nil, http.StatusNotFound)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var u User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		response(w, err.Error(), http.StatusBadRequest)
		return
	}

	Users = append(Users, u)
	response(w, Users, http.StatusOK)
}

func response(w http.ResponseWriter, body interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(body)
}
