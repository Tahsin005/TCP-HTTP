package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main () {
	addressString := "localhost:42069"
	udpAddr, err := net.ResolveUDPAddr("udp", addressString)
	if err != nil {
		fmt.Printf("Error resolving UDP address: %v\n", err)
		return
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		fmt.Printf("Error dialing UDP: %v\n", err)
		return
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(">")
		text, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error reading input: %v\n", err)
			continue
		}
		_, err = conn.Write([]byte(text))
		if err != nil {
			fmt.Printf("Error writing to UDP: %v\n", err)
			continue
		}
	}
}