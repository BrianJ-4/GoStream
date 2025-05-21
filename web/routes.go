package main

import (
	"log"
	"net/http"

	"github.com/BrianJ-4/GoStream/file"
)

func home(w http.ResponseWriter, req *http.Request) {
	log.Printf("Incoming from %s: GET /", req.RemoteAddr)
	w.Write([]byte("HOME"))
}

func video(w http.ResponseWriter, req *http.Request) {
	fileName := req.PathValue("fileName")
	log.Printf("Incoming from %s: GET /videos/%s %s", req.RemoteAddr, fileName, req.Header.Get("Range"))

	// Check if requested video exists
	err := file.CheckFileExists(fileName)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// Handle intial probe
	if req.Header.Get("Range") == "" {
		err := handleInitialProbe(w, fileName, req.RemoteAddr)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	} else { // Handle range request
		err := handleRangeRequest(w, req.Header.Get("Range"), fileName, req.RemoteAddr)
		if err != nil {
			return
		}
	}
}
