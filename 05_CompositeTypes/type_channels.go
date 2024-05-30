package main

import "fmt"

func main() {
	// Declaring an unbuffered channel
	eventsChannel := make(chan string)

	// Server-Sent Events(SSE)
	// The server sends events to the client over a single HTTP connection.
	// The client listens for events from the server and updates the UI accordingly without refreshing the page.
	//
	go sendEvents(eventsChannel)

	fmt.Println(">>> Unbuffered channel <<<")
	for event := range eventsChannel {
		fmt.Println(event)
	}

	// Declare a buffered channel
	bufferedChannel := make(chan string, 3)
	bufferedChannel <- "User registered"
	bufferedChannel <- "User logged in"
	bufferedChannel <- "User logged out"
	close(bufferedChannel)

	fmt.Println(">>> Buffered channel <<<")
	for event := range bufferedChannel {
		fmt.Println(event)
	}
}

func sendEvents(eventsChannel chan<- string) {
	eventsChannel <- "User registered"
	eventsChannel <- "User logged in"
	eventsChannel <- "User logged out"
	eventsChannel <- "User deleted"

	// Closing a channel means that no more values can be sent to it
	close(eventsChannel)
}
