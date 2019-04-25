package main

import (
	"log"
	"net/http"

	"github.com/CAIJUNYI/flickr_vis/flickr_server/controller"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir("./view")))

	http.HandleFunc("/hexagon", controller.HandleHexagon)
	http.HandleFunc("/blending", controller.HandleBlending)
	http.HandleFunc("/flickr_points", controller.HandleFlickrPoints)
	log.Fatal(http.ListenAndServe(":8888", nil))
}
