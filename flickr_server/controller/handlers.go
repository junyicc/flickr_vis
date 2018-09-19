package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"path"
	"text/template"
)

func handleHexagon(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("view", "hexagon.html"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func handleBlending(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles(path.Join("view", "blending.html"))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

func handleFlickrPoints(w http.ResponseWriter, r *http.Request) {
	log.Printf("url: %s", r.URL.Path)
	featureCollections := queryDB()
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	// w.Write([]byte(featureCollections))
	json.NewEncoder(w).Encode(featureCollections)
}