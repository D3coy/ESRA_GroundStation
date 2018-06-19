package models

import (
	"github.com/jinzhu/gorm"
)

//GPSDataSet is the schema that will contain the gps/altimeter data
type GPSDataSet struct {
	gorm.Model
	Latitude  float64
	Longitude float64
	Altitude  float64
}

// GetGPSData queries a gorm DB object and returns an instance of GPSData
func GetGPSData(db *gorm.DB) *[]GPSDataSet {
	data := new([]GPSDataSet)
	db.Find(&data)
	return data
}
