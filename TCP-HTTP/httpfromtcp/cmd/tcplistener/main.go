package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	fmt.Println("Connection accepted")

	// Read headers line by line
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Println("error reading headers:", err)
			}
			return
		}

		// Trim CRLF
		if line == "\r\n" {
			fmt.Println("(end of headers)")
			break
		}

		fmt.Print(line)
	}

	// Read body as raw bytes
	fmt.Println("Body:")
	buf := make([]byte, 1024)
	for {
		n, err := reader.Read(buf)
		if n > 0 {
			fmt.Print(string(buf[:n]))
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("error reading body:", err)
			break
		}
	}

	fmt.Println("\nConnection closed")
}

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatal("Error listening:", err)
	}
	defer listener.Close()

	fmt.Println("Listening on :42069")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn)
	}
}
