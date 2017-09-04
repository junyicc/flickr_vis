package util

import (
	"database/sql"
	"flickr_server/model"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	// Host db host
	Host = "172.21.212.167"
	// Port for remote db host
	Port = 5432
	// DBUser is user for pgsql
	DBUser = "junyi"
	// DBPasswd is password for the user above
	DBPasswd = "junyipass"
	// DBName is the database name
	DBName = "flickr"
)

func queryDB() string {
	// open and connect remote pgsql database flickr
	dbinfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Host, Port, DBUser, DBPasswd, DBName)
	db, err := sql.Open("postgres", dbinfo)
	if err != nil {
		log.Fatalf("failed to open db %s with %s", DBName, DBUser)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatalf("failed to connect db %s", DBName)
	}

	// querying
	rows, err := db.Query("SELECT image_id, image_source, taken_time, width, height, tags, image_url, owner, description, lat, lon FROM image_info")
	// rows, err := db.Query("SELECT image_url, lat, lon FROM image_info")
	if err != nil {
		log.Fatalln("failed to query image_info")
	}
	// images := make([]model.FlickrImage, 0)
	features := ""
	featureCollection := ""
	for rows.Next() {
		image := model.FlickrImage{}
		err = rows.Scan(&image.ImageID, &image.ImageSource, &image.TakenTime, &image.Width, &image.Height, &image.Tags, &image.ImageURL, &image.Owner, &image.Description, &image.Lat, &image.Lon)
		if err != nil {
			panic(err)
		}
		feature := fmt.Sprintf(`{"type": "Feature", "geometry": {"type": "Point", "coordinates": [%f, %f]}, "properties": {"ImageID": %d, "ImageSource": %q, "ImageURL": %q, "TakenTime": %q, "Owner": %q, "Width": %d, "Height": %d, "Tags": %q, "Description": %q}},`, image.Lon, image.Lat, image.ImageID, image.ImageSource, image.ImageURL, image.TakenTime.String()[:19], image.Owner, image.Width, image.Height, image.Tags, image.Description)
		features += feature
	}
	// for rows.Next() {
	// 	image := model.FlickrImage{}
	// 	err = rows.Scan(&image.ImageURL, &image.Lat, &image.Lon)
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	feature := fmt.Sprintf(`{"type": "Feature", "geometry": {"type": "Point", "coordinates": [%f, %f]}, "properties": {"ImageURL": %q}},`, image.Lon, image.Lat, image.ImageURL)
	// 	features += feature
	// }

	featureCollection = fmt.Sprintf(`{"type": "FeatureCollection", "crs": {"type": "EPSG", "properties": {"code": "4326"}}, "bbox": [-125, 25, -64, 48], "features": [%s]}`, features)
	// remove the last comma
	featureCollection = featureCollection[:len(featureCollection)-3] + featureCollection[len(featureCollection)-2:]
	return featureCollection
}
