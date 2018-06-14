package main

import (
	"github.com/d3coy/ESRA_GroundStation/gsuiserver"
)

func main() {
	// file, err := os.Create("gps_data.txt")
	// check(err)
	// defer func() {
	// 	check(file.Close())
	// }()
	// radioOut := make(chan string)
	// go communication.StartCommunications(radioOut)
	// go gsuiserver.GroundStationUI()
	// for {
	// 	log.Println(<-radioOut)
	// }
	gsuiserver.GroundStationUI()
}
