package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	SetupRoutes(r)

	fmt.Println("Running server on port 3000")
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", r))
}
