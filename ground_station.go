package main

import (
	"github.com/d3coy/ESRA_GroundStation/guiserver"
)

func main() {
	// Build Example DB
	// db := models.InitDB("./flight.db")
	// models.LoadDemo(db)

	// go communication.Start()
	guiserver.GroundStationUI()
}
