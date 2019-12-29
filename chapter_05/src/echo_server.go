package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"

	"transmission_object"
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

	var transmissionObject transmission_object.TransmissionObject
	err = json.Unmarshal(data, &transmissionObject)
	if err != nil {
		log.Panicln("Unable to unmarshal the buffer! Error: " + err.Error())
	}

	log.Println("Message = " + transmissionObject.Message)
	log.Println("Value = " + fmt.Sprintf("%f", transmissionObject.Value))

	transmissionObject.Message = "Echoed from Go: " + transmissionObject.Message
	transmissionObject.Value = 2 * transmissionObject.Value

	message, err := json.Marshal(transmissionObject)
	if err != nil {
		log.Panicln("Unable to marshal the object! Error: " + err.Error())
	}

	connection.Write(message)
}
