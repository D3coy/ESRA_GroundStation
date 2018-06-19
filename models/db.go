package models

import (
	"log"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/sqlite" // Used to open and init database
)

// InitDB takes a db file path to build the DB and automigrates schemas
func InitDB(dbPath string) *gorm.DB {
	db, err := gorm.Open("sqlite3", dbPath)
	if err != nil {
		log.Panic(err)
	}
	db.AutoMigrate(&GPSDataSet{})

	return db
}
