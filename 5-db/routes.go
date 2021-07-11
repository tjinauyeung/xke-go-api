package main

import (
	"github.com/gorilla/mux"
)

func (h *Handler) SetupRoutes(r *mux.Router) {
	r.HandleFunc("/users", h.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users", h.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", h.GetUserHandler).Methods("GET")
}
