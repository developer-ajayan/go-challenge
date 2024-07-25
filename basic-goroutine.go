package main

import (
	"fmt"
	"time"
)

// Function that will be run as a goroutine
func printNumbers() {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
		time.Sleep(500 * time.Millisecond) // Sleep for 500 milliseconds
	}
}

// Another function that will be run as a goroutine
func printLetters() {
	for _, letter := range "ABCDE" {
		fmt.Println(string(letter))
		time.Sleep(700 * time.Millisecond) // Sleep for 700 milliseconds
	}
}

func main() {
	// Running the functions as goroutines
	go printNumbers()
	go printLetters()

	// Sleep in the main goroutine to allow other goroutines to finish
	time.Sleep(4 * time.Second)
	fmt.Println("Main function finished")
}
