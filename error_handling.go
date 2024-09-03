package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// Define a custom error type for clarity
var ErrFileNotFound = errors.New("file not found")

// ReadFileWithHandling reads a file and handles potential errors
func ReadFileWithHandling(filePath string) ([]byte, error) {
  // Attempt to read the file
  data, err := ioutil.ReadFile(filePath)
  
  // Check for errors
  if err != nil {
    // Handle different error types appropriately
    if errors.Is(err, os.ErrNotExist) {
      // Handle file not found error
      return nil, ErrFileNotFound
    } else {
      // Handle other errors
      return nil, fmt.Errorf("error reading file: %w", err)
    }
  }
  
  // Return data if successful
  return data, nil
}

func main() {
  // Example usage with successful read
  data, err := ReadFileWithHandling("existing_file.txt")
  if err != nil {
    fmt.Println("Error reading file:", err)
  } else {
    fmt.Println("Successfully read file content:", string(data))
  }

  // Example usage with non-existent file
  data, err = ReadFileWithHandling("missing_file.txt")
  if err != nil {
    // Check for specific error type (ErrFileNotFound)
    if errors.Is(err, ErrFileNotFound) {
      fmt.Println("File not found:", err)
    } else {
      fmt.Println("Error reading file:", err)
    }
  } else {
    fmt.Println("This shouldn't be printed (should have error)")
  }
}