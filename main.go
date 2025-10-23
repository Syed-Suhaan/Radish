package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatalf("Failed to Start server :%v", err)
	}
	fmt.Println("Radish server started on port 6379")
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Println("Failed to accept the connection:", err)
		} else {
			fmt.Println("New client connected!")
		}
		reader := bufio.NewReader(connection)
		message, err := reader.ReadString('\n')
		if err != nil {
			log.Println("Failed to read from client:", err)
		} else {
			fmt.Printf("Received: %s", message)
			connection.Write([]byte(message))
		}
		connection.Close()
	}
}
