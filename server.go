package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	fmt.Println("New client connected!")
	reader := bufio.NewReader(conn)
	message, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Failed to read from client:", err)
	} else {
		fmt.Printf("Received: %s", message)
		conn.Write([]byte(message))
	}
	conn.Close()
}
