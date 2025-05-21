package main

import "net/http"

func handleInitialProbe(w http.ResponseWriter, req *http.Request, fileName string) error {
	// Check video file extension and set content-type
	w.Header().Set("Content-Type", "video/mp4")
	w.Header().Set("Accept-Ranges", "bytes")

	// need to fetch video size
	w.Header().Set("Content-Length", "12000")

	w.WriteHeader(http.StatusOK)
	return nil
}
