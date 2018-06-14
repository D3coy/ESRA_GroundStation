package main

import (
	"github.com/jinzhu/gorm"
)

//GPSData is the schema that will contain the gps/altimeter data
type GPSData struct {
	gorm.Model
	Latitude  float64
	Longitude float64
	Altitude  float64
}
