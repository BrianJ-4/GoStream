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
	videoName := req.PathValue("videoName")
	log.Printf("GET /videos/%s", videoName)
	w.Write([]byte(videoName))
}
