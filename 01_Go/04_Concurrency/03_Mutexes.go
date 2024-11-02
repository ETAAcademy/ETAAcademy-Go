package main

import (
	"fmt"
	"image"
	"log"
	"sync"
	"time"
)

// Visited is used to record whether a webpage has been visited
type Visited struct {
	mu      sync.Mutex
	visited map[string]int
}

// VisitLink records the access for the URL and updates the total access count for that URL
func (v *Visited) VisitLink(url string) int {
	v.mu.Lock()
	defer v.mu.Unlock()
	count := v.visited[url]
	count++
	v.visited[url] = count
	return count
}

// Long-running worker process
// func worker() {
// 	n := 0
// 	next := time.After(time.Second) // Create a timer channel
// 	for {
// 		select {
// 		case <-next: // Wait for the timer to trigger
// 			n++
// 			fmt.Println(n)
// 			next = time.After(time.Second) // Create a new timer for the next loop
// 		}
// 	}
// }
func worker() {
	pos := image.Point{X: 10, Y: 10} // Current position of the worker
	direction := image.Point{X: 1, Y: 0} // Current direction of movement
	next := time.After(time.Second) // Create a timer channel
	for {
		select {
		case <-next:
			pos = pos.Add(direction) // Update position based on direction
			fmt.Println("current position is ", pos) // Print current position
			next = time.After(time.Second) // Create a new timer for the next loop
		}
	}
}

// Command type
type command int

const (
	right = command(0) // Represents a turn to the right
	left  = command(1) // Represents a turn to the left
)

// RoverDriver is used to control the rover
type RoverDriver struct {
	commandc chan command // Channel for receiving commands
}

// NewRoverDriver creates a channel and starts the worker process
func NewRoverDriver() *RoverDriver {
	r := &RoverDriver{
		commandc: make(chan command),
	}
	go r.drive() // Start the driving process in a goroutine
	return r
}

// drive processes commands and updates the rover's position
func (r *RoverDriver) drive() {
	pos := image.Point{X: 0, Y: 0} // Initial position
	direction := image.Point{X: 1, Y: 0} // Initial direction
	updateInterval := 250 * time.Millisecond // Update interval for movement
	nextMove := time.After(updateInterval) // Create a timer for the next move
	for {
		select {
		case c := <-r.commandc: // Receive a command
			switch c {
			case right:
				// Turn right by rotating the direction
				direction = image.Point{
					X: -direction.Y,
					Y: direction.X,
				}
			case left:
				// Turn left by rotating the direction
				direction = image.Point{
					X: direction.Y,
					Y: -direction.X,
				}
			}
			log.Printf("new direction %v", direction) // Log the new direction

		case <-nextMove: // Time to move
			pos = pos.Add(direction) // Update position based on direction
			log.Printf("move to %v", pos) // Log the new position
			nextMove = time.After(updateInterval) // Reset the timer for the next move
		}
	}
}

// Left turns the rover to the left
func (r *RoverDriver) Left() {
	r.commandc <- left // Send the left command
}

// Right turns the rover to the right
func (r *RoverDriver) Right() {
	r.commandc <- right // Send the right command
}

var mu sync.Mutex // Global mutex for synchronization

func main() {
	mu.Lock() // Lock the mutex
	defer mu.Unlock() // Unlock the mutex when main function exits

	r := NewRoverDriver() // Create a new RoverDriver
	time.Sleep(3 * time.Second) // Allow rover to move for 3 seconds
	r.Left() // Turn the rover left
	time.Sleep(2 * time.Second) // Allow rover to move for 2 seconds
	r.Right() // Turn the rover right
	time.Sleep(4 * time.Second) // Allow rover to move for 4 seconds
}
