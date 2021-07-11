package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	uu := GetUsers()
	respond(w, uu, http.StatusOK)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	uid, err := uuid.Parse(params["id"])
	if err != nil {
		respond(w, err.Error(), http.StatusBadRequest)
		return
	}

	u, err := GetUser(uid)
	if err != nil {
		e, ok := errors.Cause(err).(*Error)
		if !ok {
			respond(w, "Failed to find user", http.StatusInternalServerError)
			return
		}
		if e.code == ErrorNotFound {
			respond(w, "Failed to find user", http.StatusNotFound)
			return
		}
		respond(w, "Failed to find user", http.StatusInternalServerError)
		return
	}

	respond(w, u, http.StatusOK)
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	req := new(User)
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		respond(w, err.Error(), http.StatusBadRequest)
		return
	}

	v := validator.New()
	err = v.Struct(req)
	if err != nil {
		respond(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	uu, err := CreateUser(req)
	if err != nil {
		e, ok := errors.Cause(err).(*Error)
		if !ok {
			respond(w, "Failed to create user", http.StatusInternalServerError)
			return
		}
		if e.code == ErrorNotFound {
			respond(w, "User not found", http.StatusNotFound)
			return
		}
		respond(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	respond(w, uu, http.StatusCreated)
}

func respond(w http.ResponseWriter, body interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	json.NewEncoder(w).Encode(body)
}
