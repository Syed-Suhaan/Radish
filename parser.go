package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func parseRESP(reader *bufio.Reader) ([]string, error) {

	firstByte, err := reader.ReadByte()
	if err != nil {
		return nil, err
	}

	if firstByte != '*' {
		return nil, fmt.Errorf("expected array command but got %c", firstByte)
	}

	// Read the number of elements in the array
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}

	// Parse the array length ( "*3\r\n" -> 3)
	arrayLen, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		return nil, fmt.Errorf("invalid array length: %s", line)
	}

	// Create a slice to hold all the command arguments
	args := make([]string, arrayLen)

	for i := 0; i < arrayLen; i++ {
		// Each argument must be a Bulk String (starting with '$')
		argType, err := reader.ReadByte()
		if err != nil {
			return nil, err
		}
		if argType != '$' {
			return nil, fmt.Errorf("expected bulk string but got %c", argType)
		}

		// Read the length of the bulk string
		line, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		// Parse the string length ("$4\r\n" -> 4)
		strLen, err := strconv.Atoi(strings.TrimSpace(line))
		if err != nil {
			return nil, fmt.Errorf("invalid bulk string length: %s", line)
		}

		// Read the string data itself, plus the trailing \r\n
		data := make([]byte, strLen+2) // +2 for \r\n
		if _, err := io.ReadFull(reader, data); err != nil {
			return nil, err
		}

		// Store the string, without the trailing \r\n
		args[i] = string(data[:strLen])
	}

	return args, nil
}
