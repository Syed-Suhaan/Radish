package main

import (
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
		connection.Close()
	}
}
