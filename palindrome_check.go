package main

import "fmt"

func isPalindrome(word string) bool {
    length := len(word)
    for i := 0; i < length/2; i++ {
        if word[i] != word[length-1-i] {
            return false
        }
    }
    return true
}

func main() {
    word := "malayalam"
    if isPalindrome(word) {
        fmt.Println("Number is palindrome")
    } else {
        fmt.Println("Number is not palindrome")
    }
}