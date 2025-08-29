package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	ch := make(chan string)
	go func() {
		defer f.Close()
		defer close(ch)

		buffer := make([]byte, 8)
		var currentLine string

		for {
			bytesRead, err := f.Read(buffer)
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
	fileName := "messages.txt"

	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	// defer file.Close()

	// buffer := make([]byte, 8)
    // var currentLine string

	// for {
	// 	bytesRead, err := file.Read(buffer)

	// 	if err == io.EOF {
	// 		if bytesRead > 0 {
	// 			parts := strings.Split(string(buffer[:bytesRead]), "\n")
	// 			currentLine += parts[0]
	// 			if currentLine != "" {
	// 				fmt.Printf("read: %s\n", currentLine)
	// 			}
	// 			if len(parts) > 1 {
	// 				for _, part := range parts[1:] {
	// 					if part != "" {
	// 						fmt.Printf("read: %s\n", part)
	// 					}
	// 				}
	// 			}
	// 		}
    //         fmt.Println("read: end")
	// 		break
	// 	}

	// 	if bytesRead > 0 {
	// 		parts := strings.Split(string(buffer[:bytesRead]), "\n")
	// 		for i, part := range parts {
	// 			if i < len(parts) - 1 {
	// 				currentLine += part
	// 				if currentLine != "" {
	// 					fmt.Printf("read: %s\n", currentLine)
	// 					currentLine = ""
	// 				}
	// 			} else {
	// 				currentLine += part
	// 			}
	// 		}
	// 	}
	// }

    lines := getLinesChannel(file)
	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}
}
