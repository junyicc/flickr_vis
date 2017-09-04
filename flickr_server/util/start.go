package util

import (
	"log"
	"net/http"
)

// Start server
func Start() {
	http.HandleFunc("/flickr_points", handleFlickrPoints)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
