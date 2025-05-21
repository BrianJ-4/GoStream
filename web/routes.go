package main

import (
	"log"
	"net/http"

	"github.com/BrianJ-4/GoStream/file"
)

func home(w http.ResponseWriter, req *http.Request) {
	log.Printf("GET /")
	w.Write([]byte("HOME"))
}

func video(w http.ResponseWriter, req *http.Request) {
	fileName := req.PathValue("fileName")
	log.Printf("GET /videos/%s", fileName)

	// Check if requested video exists
	err := file.CheckFileExists(fileName)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Check for intial probe
	if req.Header.Get("Range") == "" {
		err := handleInitialProbe(w, fileName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
}
