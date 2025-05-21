package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/BrianJ-4/GoStream/file"
)

var extensionMapping = map[string]string{
	".mp4":  "mp4",
	".webm": "webm",
	".ogg":  "ogg",
}

const chunkSize = int64(1024 * 1024) // 1mb

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
		log.Print("Error getting video size: ", err)
		return err
	}
	w.Header().Set("Content-Length", strconv.FormatInt(size, 10))

	// Add Accept-Ranges header
	w.Header().Set("Accept-Ranges", "bytes")

	w.WriteHeader(http.StatusOK)
	return nil
}

func handleRangeRequest(w http.ResponseWriter, requestRange string, fileName string) error {
	// Open Video
	video, err := file.OpenFile(fileName)
	if err != nil {
		log.Print("Error opening file: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}
	defer video.Close()

	// Get video size
	size, err := file.GetFileSize(video)
	if err != nil {
		log.Print("Error getting video size: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	r, err := parseRange(requestRange, size)
	if err != nil {
		w.WriteHeader(http.StatusRequestedRangeNotSatisfiable)
		return err
	}
	if 0 > 1 {
		fmt.Print(r)
	}

	// Add Accept-Ranges header
	w.Header().Set("Accept-Ranges", "bytes")

	w.WriteHeader(http.StatusPartialContent)
	return nil
}

type Range struct {
	Start  int64
	Length int64
}

func parseRange(requestRange string, size int64) (Range, error) {
	// Range examples: bytes=0-; bytes=524288-; bytes=500-999; bytes=-100
	var r Range

	// Not supporting multipart ranges
	if strings.Contains(requestRange, ",") {
		err := errors.New("multipart ranges not supported")
		log.Print("Error parsing range: ", err, ": ", requestRange)
		return r, err
	}

	requestRange = strings.Trim(requestRange, "bytes=") // Strips "bytes="
	parts := strings.Split(requestRange, "-")

	if len(parts) != 2 {
		err := errors.New("invalid range format")
		log.Print("Error parsing range: ", err, ": bytes=", requestRange)
		return r, err
	}

	// Suffix range
	if parts[0] == "" {
		length, err := strconv.ParseInt(parts[1], 10, 64)
		if err != nil {
			log.Print("Error parsing suffix range: ", err, ": bytes=", requestRange)
			return r, err
		}
		if length <= 0 {
			err := errors.New("invalid suffix range: length must be greater than 0")
			log.Print("Error parsing suffix range: ", err, ": bytes=", requestRange)
			return r, err
		}
		// If last-byte-pos is greater than or equal to the current length
		// of the representation data, the byte range is interpreted as the
		// remainder of the representation
		if length > size {
			length = size
		}
		r.Start = size - length
		r.Length = length
	} else {
		start, err := strconv.ParseInt(parts[0], 10, 64)
		if err != nil {
			log.Print("Error parsing range: ", err, ": bytes=", requestRange)
			return r, err
		}
		var length int64
		// Normal range
		if parts[1] != "" {
			end, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				log.Print("Error parsing range: ", err, ": bytes=", requestRange)
				return r, err
			}
			if end >= size {
				end = size - 1
			}
			length = end - start + 1
		} else { // Prefix range
			length = size - start
		}
		r.Start = start
		r.Length = length
	}

	err := r.validateRange(size)
	if err != nil {
		log.Print("Error parsing range: ", err, ": bytes=", requestRange)
		return r, err
	}

	// Limit size of data to chunkSize (1mb)
	if r.Length > chunkSize {
		r.Length = chunkSize
	}

	fmt.Println(r)
	return r, nil
}

func (r *Range) validateRange(size int64) error {
	if r.Start >= size {
		return errors.New("range start beyond file size")
	}
	if r.Length <= 0 {
		return errors.New("range length must be positive")
	}
	return nil
}
