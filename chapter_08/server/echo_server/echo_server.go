package main

import (
	"fmt"
	"log"
	"net"
	"strings"

	"transceiver"
	"transmission_object"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type EchoServer struct{}

func (es *EchoServer) Echo(context context.Context, request *transceiver.EchoRequest) (*transceiver.EchoResponse, error) {
	log.Println("Message = " + (*request).FromClient.GetMessage())
	log.Println("Value = " + fmt.Sprintf("%f", (*request).FromClient.GetValue()))
	server_message := "Received from client: " +
		(*request).FromClient.GetMessage()
	server_value := (*request).FromClient.Value * 2
	from_server := transmission_object.TransmissionObject{
		Message: server_message,
		Value:   server_value,
	}
	return &transceiver.EchoResponse{
		FromServer: &from_server,
	}, nil
}

func (es *EchoServer) UpperCase(contest context.Context, request *transceiver.UpperCaseRequest) (*transceiver.UpperCaseResponse, error) {
	log.Println("Original = " + (*request).GetOriginal())
	return &transceiver.UpperCaseResponse{
		UpperCased: strings.ToUpper((*request).GetOriginal()),
	}, nil
}

func main() {
	log.Println("Spinning up the Echo Server in Go...")
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Panicln("Unable to listen: " + err.Error())
	}
	defer listen.Close()
	defer log.Println("Connection now closed.")

	grpc_server := grpc.NewServer()
	transceiver.RegisterTransceiverServer(grpc_server, &EchoServer{})
	err = grpc_server.Serve(listen)
	if err != nil {
		log.Panicln("Unable to start serving! Error: " + err.Error())
	}
}
