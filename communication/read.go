// Package communication contains radio communcation programs to run
package communication

import (
	"log"
	"time"

	"github.com/tarm/serial"
)

const (
	serialPort = "/dev/ttyUSB0"
	baudRate   = 57600
	timeOut    = time.Second * 60
)

// readLines begins connection to a radio via a serial port defined in constants.
// It interprets data recieved from the transmitter and records the data into a sql server.
func readLines(in <-chan byte, out chan<- string) {
	sentence := ""
	var b byte
	for {
		b = <-in
		sentence += string(b)
		if b == byte(10) {
			out <- sentence
			sentence = ""
		}
	}
}

// readSerial will start reading bytes from the Serial Port and feed them thorugh the byte chan given on call.
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
