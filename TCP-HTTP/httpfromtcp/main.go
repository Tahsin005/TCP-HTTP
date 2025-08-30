package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func getLinesChannel(conn net.Conn) <-chan string {
	ch := make(chan string)
	go func() {
		defer conn.Close()
		defer close(ch)

		buffer := make([]byte, 8)
		var currentLine string

		for {
			bytesRead, err := conn.Read(buffer)
			if err == io.EOF {
				if bytesRead > 0 {
					parts := strings.Split(string(buffer[:bytesRead]), "\n")
					currentLine += parts[0]
					if currentLine != "" {
						ch <- currentLine
					}
					if len(parts) > 1 {
						for _, part := range parts[1:] {
							if part != "" {
								ch <- part
							}
						}
					}
				}
                fmt.Println("read: end")
				return
			}

			if err != nil {
				log.Printf("error reading file: %v", err)
				return
			}

			if bytesRead > 0 {
				parts := strings.Split(string(buffer[:bytesRead]), "\n")
				for i, part := range parts {
					if i < len(parts) - 1 {
						currentLine += part
						if currentLine != "" {
							ch <- currentLine
							currentLine = ""
						}
					} else {
						currentLine += part
					}
				}
			}
		}
	}()
	return ch
}

func main() {
	listener, err := net.Listen("tcp", ":42069")
    if err != nil {
		fmt.Println("Error listening:", err.Error())
		return
	}
	defer listener.Close()

	fmt.Println("Listening on :42069")

    for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			continue
		}

		fmt.Println("Connection accepted")
		lines := getLinesChannel(conn)
		for line := range lines {
			fmt.Println(line)
		}
		fmt.Println("Connection closed")
	}
}
