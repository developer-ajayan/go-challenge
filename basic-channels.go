package main

import (
	"fmt"
)

func main() {
	ch := make(chan int) // Creating a channel of type int

	go func() {
		ch <- 42 // Sending data into the channel (blocks until the data is received)
		fmt.Println("Sent 42 to the channel")
	}()

	value := <-ch // Receiving data from the channel (blocks until data is available)
	fmt.Println("Received:", value)
}
