package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, req *http.Request) {
	log.Printf("GET /")
	w.Write([]byte("HOME"))
}

func video(w http.ResponseWriter, req *http.Request) {
	fileName := req.PathValue("fileName")
	log.Printf("GET /videos/%s", fileName)

	// Will need to check if video exists

	// Check for intial probe
	if req.Header.Get("Range") == "" {
		handleInitialProbe(w, req, fileName)
		println("HIT")
	}

	//w.Write([]byte(fileName))
}
