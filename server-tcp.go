package main

import (
	"log"
	"time"

	"github.com/tbrandon/mbserver"
)

func main() {
	serv := mbserver.NewServer()
	err := serv.ListenTCP("127.0.0.1:5020")
	if err != nil {
		log.Printf("%v\n", err)
	}
	defer serv.Close()

	// Wait forever
	for {
		time.Sleep(1 * time.Second)
	}
}
