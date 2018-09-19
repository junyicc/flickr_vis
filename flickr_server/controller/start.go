package controller

import (
	"log"
	"net/http"
	"path"
)

// Start server
func StartServer() {
	http.Handle("/lib/", http.StripPrefix("/lib/", http.FileServer(http.Dir(path.Join("view", "lib")))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir(path.Join("view", "js")))))

	http.HandleFunc("/hexagon", handleHexagon)
	http.HandleFunc("/blending", handleBlending)
	http.HandleFunc("/flickr_points", handleFlickrPoints)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
