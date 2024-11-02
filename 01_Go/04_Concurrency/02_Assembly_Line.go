package main

import (
	"fmt"
	"strings"
)

// Source function that sends messages to the downstream channel
func sourceGopher(downStream chan string) {
	for _, v := range []string{"hello world", "a bad apple", "goodbye all"} {
		downStream <- v // Send each message to the downstream channel
	}
	// downStream <- "" // Optionally send an empty string to signal completion
	close(downStream) // Close the downstream channel after sending all messages
}

// Middle function that filters messages
func filterGopher(upStream, downStream chan string) {
	// for {
	// 	// item := <-upStream
	// 	item, ok := <-upStream // Receive from the upstream channel
	// 	if !ok {
	// 		// downStream <- "" // Optionally send an empty string to signal completion
	// 		close(downStream) // Close the downstream channel when done
	// 		return
	// 	}
	// 	if !strings.Contains(item, "bad") {
	// 		downStream <- item // Send item to downstream if it doesn't contain "bad"
	// 	}
	// }
	// Use range to read values from the channel
	for item := range upStream {
		if !strings.Contains(item, "bad") {
			downStream <- item // Send item to downstream if it doesn't contain "bad"
		}
	}
	close(downStream) // Close the downstream channel after processing
}

// Downstream function that prints messages
func printGopher(upStream chan string) {
	// for {
	// 	v := <-upStream // Receive from the upstream channel
	// 	if v == "" {
	// 		return // Exit if an empty string is received
	// 	}
	// 	fmt.Println(v) // Print the received message
	// }
	for v := range upStream {
		fmt.Println(v) // Print each message received from the upstream channel
	}
}

func main() {
	c0 := make(chan string) // Create the upstream channel
	c1 := make(chan string) // Create the downstream channel
	go sourceGopher(c0)    // Start the source goroutine
	go filterGopher(c0, c1) // Start the filter goroutine
	printGopher(c1)        // Start printing from the downstream channel
}
