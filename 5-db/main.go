package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Handler struct {
	DB *gorm.DB
}

func main() {
	r := mux.NewRouter()

	h := Handler{
		DB: SetupDB(),
	}
	h.SetupRoutes(r)

	fmt.Println("Running server on port 3000")
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", r))
}
