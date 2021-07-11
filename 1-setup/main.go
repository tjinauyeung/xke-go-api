package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	r := http.NewServeMux()

	r.HandleFunc("/health", healthCheck)

	fmt.Printf("Running server on port 3000")
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", r))
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "OK")
}
