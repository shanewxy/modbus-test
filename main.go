package main

import (
	"github.com/goburrow/modbus"
	"log"
	"os"
	"time"
)

func main() {
	//testTCP()
	testRTU()
}

func testTCP() {
	// Modbus TCP
	handler := modbus.NewTCPClientHandler("localhost:5020")
	handler.Timeout = 10 * time.Second
	handler.SlaveId = 0x01
	handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)
	// Connect manually so that multiple requests are handled in one connection session
	err := handler.Connect()
	defer handler.Close()

	client := modbus.NewClient(handler)
	results, err := client.ReadHoldingRegisters(8, 1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(results)
	results, err = client.WriteSingleRegister(1, 8)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(results)
	results, err = client.WriteMultipleCoils(1, 10, []byte{4, 3})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(results)
}

func testRTU() {
	// Modbus RTU/ASCII
	handler := modbus.NewRTUClientHandler("/dev/ttys010")
	handler.BaudRate = 115200
	handler.DataBits = 8
	handler.Parity = "N"
	handler.StopBits = 1
	handler.SlaveId = 1
	handler.Timeout = 5 * time.Second

	err := handler.Connect()
	if err != nil {
		log.Fatal("error connecting: ", err)
	}
	defer handler.Close()

	client := modbus.NewClient(handler)
	results, err := client.ReadHoldingRegisters(8, 1)
	if err != nil {
		log.Fatal("error reading: ", err)
	}
	log.Println(results)
}
