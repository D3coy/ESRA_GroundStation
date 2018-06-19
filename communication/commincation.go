package communication

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/d3coy/ESRA_GroundStation/models"
)

const dbPath = "./flight.db"

// Start begins the reading of data from serial io and formats it
func Start() {
	db := models.InitDB(dbPath)
	rawBytes := make(chan byte)
	rawStrings := make(chan string)
	go readSerial(rawBytes)
	go readLines(rawBytes, rawStrings)
	for {
		data := <-rawStrings
		ds := new(models.GPSDataSet)
		err := json.NewDecoder(strings.NewReader(data)).Decode(&ds)
		if err != nil {
			log.Println("Failed to Decode json", err)
			continue
		}
		db.Create(&ds)
	}
}
