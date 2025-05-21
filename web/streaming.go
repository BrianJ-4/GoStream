package main

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/BrianJ-4/GoStream/file"
)

var extensionMapping = map[string]string{
	".mp4":  "mp4",
	".webm": "webm",
	".ogg":  "ogg",
}

func handleInitialProbe(w http.ResponseWriter, fileName string) error {
	// Open Video
	video, err := file.OpenFile(fileName)
	if err != nil {
		log.Print("Error opening file: ", err)
		return err
	}
	defer video.Close()

	// Check video file extension and set content-type
	ext := file.GetFileExtension(video)
	elem, ok := extensionMapping[ext]
	if !ok {
		err := errors.New("unsupported file type")
		log.Print("Error setting Content-Type: ", err)
		return err
	}
	w.Header().Set("Content-Type", "video/"+elem)

	// Set Content-Length
	size, err := file.GetFileSize(video)
	if err != nil {
		log.Print("Error getting video length: ", err)
		return err
	}
	w.Header().Set("Content-Length", strconv.FormatInt(size, 10))

	// Add Accept-Ranges header
	w.Header().Set("Accept-Ranges", "bytes")

	w.WriteHeader(http.StatusOK)
	return nil
}
