package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fileName := "messages.txt"

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	buffer := make([]byte, 8)

	for {
		bytesRead, err := file.Read(buffer)

		if err == io.EOF {
			if bytesRead > 0 {
				fmt.Printf("%s\n", buffer[:bytesRead])
			}
			break
		}

		fmt.Printf("read: %s\n", buffer[:bytesRead])
	}
}
