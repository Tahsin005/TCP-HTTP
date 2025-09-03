package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func handleConnection(conn net.Conn) <-chan string {
	out := make(chan string)

	go func() {
		defer conn.Close()
		defer close(out)

		reader := bufio.NewReader(conn)
		out <- "Connection accepted"

		// Read headers line by line
		for {
			line, err := reader.ReadString('\n')
			if err != nil {
				if err != io.EOF {
					log.Println("error reading headers:", err)
				}
				return
			}

			// Detect end of headers
			if line == "\r\n" {
				out <- "(end of headers)"
				break
			}

			out <- line[:len(line)-1] // trim trailing \n
		}

		// Read body as raw bytes
		out <- "Body:"
		buf := make([]byte, 1024)
		for {
			n, err := reader.Read(buf)
			if n > 0 {
				out <- string(buf[:n])
			}
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Println("error reading body:", err)
				break
			}
		}

		out <- "Connection closed"
	}()

	return out
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

		lines := handleConnection(conn)
		for line := range lines {
			fmt.Println(line)
		}
	}
}
