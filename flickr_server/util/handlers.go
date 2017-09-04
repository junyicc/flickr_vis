package util

import (
	"encoding/json"
	"log"
	"net/http"
)

func handleFlickrPoints(w http.ResponseWriter, r *http.Request) {
	log.Printf("url: %s", r.URL.Path)
	featureCollections := queryDB()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	// w.Write([]byte(featureCollections))
	json.NewEncoder(w).Encode(featureCollections)
}
