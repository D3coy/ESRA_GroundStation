// Package communication contains radio communcation programs to run
package communication

import (
	"log"
	"time"

	nmea "github.com/adrianmo/go-nmea"
	"github.com/tarm/serial"
)

const (
	serialPort = "COM4"
	baudRate   = 57600
	timeOut    = time.Second * 60
)

// StartCommunications begins connection to a radio via a
// serial port defined in constants. It interprets data recieved
// from the transmitter and records the data into a sql server
func StartCommunications(out chan<- string) {
	log.Println("test world")
	data := make(chan byte)
	sentence := ""
	var b byte
	go readSerial(data)

	for {
		b = <-data
		sentence += string(b)
		if b == byte(10) {
			s, err := nmea.Parse(sentence)
			if err != nil {
				log.Println("Failed to Parse Sentence: ", err)
				continue
			}
			m := s.(nmea.GPRMC)
			log.Printf("Raw sentence: %v\n", m)
			log.Printf("Time: %s\n", m.Time)
			log.Printf("Validity: %s\n", m.Validity)
			log.Printf("Latitude GPS: %s\n", nmea.FormatGPS(m.Latitude))
			log.Printf("Latitude DMS: %s\n", nmea.FormatDMS(m.Latitude))
			log.Printf("Longitude GPS: %s\n", nmea.FormatGPS(m.Longitude))
			log.Printf("Longitude DMS: %s\n", nmea.FormatDMS(m.Longitude))
			log.Printf("Speed: %f\n", m.Speed)
			log.Printf("Course: %f\n", m.Course)
			log.Printf("Date: %s\n", m.Date)
			log.Printf("Variation: %f\n", m.Variation)
			out <- "Successfully processed Sentence"
			sentence = ""
		}
	}
}

// readSerial will start reading bytes from the Serial
// Port and feed them thorugh the byte chan given on call.
func readSerial(data chan<- byte) {
	log.Printf("Opening Serial Port: %v\n", serialPort)
	c := &serial.Config{Name: serialPort, Baud: baudRate, ReadTimeout: timeOut}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	buf := make([]byte, 128)
	for {
		n, err := s.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		for _, b := range buf[:n] {
			data <- b
		}
	}
}
