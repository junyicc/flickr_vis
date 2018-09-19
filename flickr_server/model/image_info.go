package model

import "time"

// FlickrImage struct for pgsql
type FlickrImage struct {
	ImageID     int
	ImageSource string
	TakenTime   time.Time
	Width       int
	Height      int
	Tags        string
	ImageURL    string
	Owner       string
	Description string
	Lat         float64
	Lon         float64
}
