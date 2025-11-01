package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func handleConnection(conn net.Conn) {
	fmt.Println("New client connected!")
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		args, err := parseRESP(reader)
		if err != nil {
			log.Println("Client disconnected or read error:", err)
			break
		}
		if len(args) == 0 {
			continue
		}
		command := strings.ToUpper(args[0])
		if command == "PING" {
			conn.Write([]byte("+PONG\r\n"))
		} else {
			conn.Write([]byte("-ERR unknown command\r\n"))
		}
	}
}
