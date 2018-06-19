package models

import "github.com/jinzhu/gorm"

// LoadDemo preloads a db with points to demo the UI.
func LoadDemo(db *gorm.DB) {
	points := []*GPSDataSet{
		&GPSDataSet{Latitude: 3229.5, Longitude: -10676.0, Altitude: 0},
		&GPSDataSet{Latitude: 3229.8, Longitude: -10675.7, Altitude: 4225},
		&GPSDataSet{Latitude: 3230.1, Longitude: -10675.4, Altitude: 7825},
		&GPSDataSet{Latitude: 3230.4, Longitude: -10675.1, Altitude: 10850},
		&GPSDataSet{Latitude: 3230.7, Longitude: -10674.8, Altitude: 13350},
		&GPSDataSet{Latitude: 3231.0, Longitude: -10674.5, Altitude: 15375},
		&GPSDataSet{Latitude: 3231.3, Longitude: -10674.2, Altitude: 16975},
		&GPSDataSet{Latitude: 3231.6, Longitude: -10673.9, Altitude: 18200},
		&GPSDataSet{Latitude: 3231.9, Longitude: -10673.6, Altitude: 19100},
		&GPSDataSet{Latitude: 3232.2, Longitude: -10673.3, Altitude: 19725},
		&GPSDataSet{Latitude: 3232.5, Longitude: -10673.0, Altitude: 19725},
		&GPSDataSet{Latitude: 3232.8, Longitude: -10672.7, Altitude: 19100},
		&GPSDataSet{Latitude: 3233.1, Longitude: -10672.4, Altitude: 18200},
		&GPSDataSet{Latitude: 3233.4, Longitude: -10672.1, Altitude: 16975},
		&GPSDataSet{Latitude: 3233.7, Longitude: -10671.8, Altitude: 15375},
		&GPSDataSet{Latitude: 3234.0, Longitude: -10671.5, Altitude: 13350},
		&GPSDataSet{Latitude: 3234.3, Longitude: -10671.2, Altitude: 10850},
		&GPSDataSet{Latitude: 3234.6, Longitude: -10670.9, Altitude: 7825},
		&GPSDataSet{Latitude: 3234.9, Longitude: -10670.6, Altitude: 4225},
		&GPSDataSet{Latitude: 3235.2, Longitude: -10670.3, Altitude: 0},
	}
	for _, ds := range points {
		db.Create(ds)
	}
}
