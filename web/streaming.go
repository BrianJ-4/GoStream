package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/BrianJ-4/GoStream/file"
)

func handleInitialProbe(w http.ResponseWriter, fileName string) error {
	// Open Video
	video, err := file.OpenFile(fileName)
	if err != nil {
		log.Print(err)
		return err
	}

	// Check video file extension and set content-type
	w.Header().Set("Content-Type", "video/mp4")
	w.Header().Set("Accept-Ranges", "bytes")

	// Set Content-Length
	size, err := file.GetFileLength(video)
	if err != nil {
		log.Print(err)
		return err
	}
	w.Header().Set("Content-Length", strconv.FormatInt(size, 10))

	w.WriteHeader(http.StatusOK)
	return nil
}
