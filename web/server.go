package main

import (
	"log"
	"net/http"
)

var port = "8090"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /videos/{fileName}", video)

	log.Printf("Starting server on :%s", port)
	err := http.ListenAndServe(":"+port, mux)
	log.Fatal(err)
}
