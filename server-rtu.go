package main

import (
	"github.com/goburrow/serial"
	"log"
	"time"

	"github.com/tbrandon/mbserver"
)

func main() {

	serv := mbserver.NewServer()
	// socat -d -d pty,raw,echo=0 pty,raw,echo=0
	err := serv.ListenRTU(&serial.Config{Address: "/dev/ttys008"})
	if err != nil {
		log.Printf("%v\n", err)
	}
	defer serv.Close()

	// Wait forever
	for {
		time.Sleep(1 * time.Second)
	}
}
