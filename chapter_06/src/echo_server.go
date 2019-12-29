package main

import (
	"fmt"
	"log"
	"net"

	"transmission_object"

	"github.com/golang/protobuf/proto"
)

func main() {
	log.Println("Spinning up the Echo Server in Go...")

	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Panicln("Unable to listen: " + err.Error())
	}
	defer listen.Close()

	connection, err := listen.Accept()
	if err != nil {
		log.Panicln("Cannot accept a connection! Error: " + err.Error())
	}

	log.Println("Receiving on a new connection")
	defer connection.Close()
	defer log.Println("Connection now closed.")

	buffer := make([]byte, 2048)
	size, err := connection.Read(buffer)
	if err != nil {
		log.Println("Cannot read from the buffer! Error: " + err.Error())
	}

	data := buffer[:size]

	transmissionObject := &transmission_object.TransmissionObject{}
	err = proto.Unmarshal(data, transmissionObject)
	if err != nil {
		log.Panicln("Unable to unmarshal the buffer! Error: " + err.Error())
	}

	log.Println("Message = " + transmissionObject.GetMessage())
	log.Println("Value = " + fmt.Sprintf("%f", transmissionObject.GetValue()))

	transmissionObject.Message = "Echoed from Go: " + transmissionObject.GetMessage()
	transmissionObject.Value = 2 * transmissionObject.GetValue()

	message, err := proto.Marshal(transmissionObject)
	if err != nil {
		log.Panicln("Unable to marshal the object! Error: " + err.Error())
	}

	connection.Write(message)
}
