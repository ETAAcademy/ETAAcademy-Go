package main

import (
	"fmt"
	"time"
)

func sleepyGopher() {
	time.Sleep(time.Second * 3) // Sleep for 3 seconds to simulate task processing
	fmt.Println("...snore...")
}

func sleepyGopher2(id int) {
	time.Sleep(time.Second * 3) // Sleep for 3 seconds to simulate task processing
	fmt.Println("...snore...", id)
}

func main() {	
	go sleepyGopher() // Create a new goroutine with the 'go' keyword
	fmt.Println("this is main func")
	time.Sleep(time.Second * 4) // If this line is commented out, the main function won't wait for the goroutine to finish, and "...snore..." won't be printed.

	// Start 5 goroutines in a loop
	for i := 0; i < 5; i++ {
		go sleepyGopher()
	}
	time.Sleep(time.Second * 4) // Wait for 4 seconds to let the goroutines finish

	for i := 0; i < 5; i++ {
		go sleepyGopher2(i) // Start 5 more goroutines, passing the index as an argument
	}
	time.Sleep(time.Second * 4) // Wait for 4 seconds to let the goroutines finish
}

// 2. Channel

// func sleepyGopher(id int, c chan int) {
// 	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond) // Sleep for a random duration
// 	fmt.Println("...snore...", id)
// 	c <- id // Send the gopher's ID back through the channel
// }

// func main() {
// 	c := make(chan int) // Create a channel
// 	for i := 0; i < 5; i++ {
// 		go sleepyGopher(i, c) // Start 5 goroutines, each sending their ID through the channel
// 	}
// 	for i := 0; i < 5; i++ {
// 		gopherID := <-c // Receive the gopher's ID from the channel
// 		fmt.Println("gopher", gopherID, "has finished sleeping")
// 	}

// 	// timeout := time.After(2 * time.Second)
// 	// for i := 0; i < 5; i++ {
// 	// select {
// 	// case gopherID := <-c: // Wait for a gopher to finish
// 	// 	fmt.Println("gopher", gopherID, "has finished sleeping")
// 	// case <-timeout: // Wait until timeout
// 	// 	fmt.Println("my patience ran out")
// 	// 	return
// 	// }
// }

// 3. Deadlock

// func main(){
// 	c := make(chan int) // Create a channel with a specified type
// 	go func() {c <- 2}() // Start a goroutine that sends 2 to the channel
// 	<- c // Receive from the channel
// }
