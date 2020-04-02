package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/goburrow/modbus"
)

func main() {
	//testTCP()
	testRTU()
	//data := make([]byte, 8)
	//f, err := strconv.ParseFloat("3.3", 64)
	//if err != nil {
	//}
	//binary.BigEndian.PutUint64(data, math.Float64bits(f))
	//fmt.Println(data)
}
func testTCP() {
	// Modbus TCP
	handler := modbus.NewTCPClientHandler("localhost:502")
	handler.Timeout = 10 * time.Second
	handler.SlaveId = 0x01
	handler.Logger = log.New(os.Stdout, "test: ", log.LstdFlags)
	// Connect manually so that multiple requests are handled in one connection session
	err := handler.Connect()
	defer handler.Close()

	client := modbus.NewClient(handler)
	//read, err := client.ReadHoldingRegisters(0, 10)
	//read, err := client.ReadInputRegisters(1, 3)
	if err != nil {
		log.Fatal(err)
	}
	//log.Println(results)
	//results, err = client.WriteSingleRegister(1, 8)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//log.Println(results)
	var test = "true"
	var bt = []byte(test)
	fmt.Println(bt)
	//if len(bt)%2 != 0 {
	//	bt = append(bt, 0)
	//}
	l := uint16(len(bt))
	fmt.Println(l)
	results, err := client.WriteMultipleCoils(0, 9, []byte{10})
	//results, err := client.WriteMultipleCoils(5, 2, bt)
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, 9)
	fmt.Println(bs)
	//results, err := client.WriteMultipleRegisters(0, 1, []byte{1,0})
	//results, err := client.WriteMultipleRegisters(1, l/2, bt)
	fmt.Println(string([]byte{4, 3}))
	if err != nil {
		log.Fatal(err)
	}
	log.Println(results)

	//read, err := client.ReadCoils(1, 8)
	//read, err := client.ReadHoldingRegisters(2, 3)
	//log.Println(read)
	//log.Println(string(read))
}

func testRTU() {
	// Modbus RTU/ASCII
	handler := modbus.NewRTUClientHandler("/dev/ttyUSB0")
	handler.BaudRate = 9600
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
	results, err := client.ReadHoldingRegisters(0, 1)
	if err != nil {
		log.Fatal("error reading: ", err)
	}
	log.Println(results)
}
