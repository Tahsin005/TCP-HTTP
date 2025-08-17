// Chirpy is a social network similar to Twitter.

// One of Chirpy's servers is processing requests unbelievably slowly. Use a goroutine to fix the bug in the handleRequests (not handleRequest) function. The server should be able to process all the requests within the time limit.

package main

import (
	"fmt"
	"time"
)

func handleRequests(reqs <-chan request) {
	for req := range reqs {
		go handleRequest(req)
	}
}

type request struct {
	path string
}

func main() {
	reqs := make(chan request, 100)
	go handleRequests(reqs)
	for
 i := 0; i < 4; i++ {
		reqs <- request{path: fmt.Sprintf("/path/%d", i)}
		time.Sleep(500 * time.Millisecond)
	}

	time.Sleep(5 * time.Second)
	fmt.Println("5 seconds passed, killing server")
}

func handleRequest(req request) {
	fmt.Println("Handling request for", req.path)
	time.Sleep(2 * time.Second)

	fmt.Println("Done with request for", req.path)
}
