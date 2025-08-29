package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fileName := "messages.txt"

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer file.Close()

	buffer := make([]byte, 8)
    var currentLine string

	for {
		bytesRead, err := file.Read(buffer)

		if err == io.EOF {
			if bytesRead > 0 {
				parts := strings.Split(string(buffer[:bytesRead]), "\n")
				currentLine += parts[0]
				if currentLine != "" {
					fmt.Printf("read: %s\n", currentLine)
				}
				if len(parts) > 1 {
					for _, part := range parts[1:] {
						if part != "" {
							fmt.Printf("read: %s\n", part)
						}
					}
				}
			}
            fmt.Println("read: end")
			break
		}

		if bytesRead > 0 {
			parts := strings.Split(string(buffer[:bytesRead]), "\n")
			for i, part := range parts {
				if i < len(parts) - 1 {
					currentLine += part
					if currentLine != "" {
						fmt.Printf("read: %s\n", currentLine)
						currentLine = ""
					}
				} else {
					currentLine += part
				}
			}
		}
	}
}
