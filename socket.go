package main

import (
	"fmt"
	"net"
)

func createSocket(port string) {
	incomingConnection, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Socket started at port %s", port)

	for {
		connection, err := incomingConnection.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConnection(connection)
	}
}

func handleConnection(connection net.Conn) {
	defer connection.Close()

	// read incoming data
	buffer := make([]byte, 1024)
	_, err := connection.Read(buffer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Received: %s", buffer)
}
