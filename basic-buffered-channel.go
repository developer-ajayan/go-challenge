package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3) // Create a buffered channel with a capacity of 3

	// Goroutine to send data into the channel
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Sending %d to channel\n", i)
			ch <- i // Send value to the channel
			fmt.Printf("Sent %d to channel\n", i)
			time.Sleep(500 * time.Millisecond) // Simulate work
		}
		close(ch) // Close the channel after sending all values
	}()

	// Goroutine to receive data from the channel
	go func() {
		for value := range ch {
			fmt.Printf("Received %d from channel\n", value)
			time.Sleep(700 * time.Millisecond) // Simulate work
		}
	}()

	// Sleep in the main goroutine to allow other goroutines to finish
	time.Sleep(5 * time.Second)
	fmt.Println("Main function finished")
}
