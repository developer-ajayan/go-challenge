package main

import "fmt"

func main() {
	var word string
	word = "ajayan"
	var reverse_of_word string
	
	// runes := []rune(word)
	// length := len(runes)
	// for i := 0; i < length/2; i++ {
	// 	runes[i], runes[(length-1)-i] = runes[(length-1)-i], runes[i]
	// }
	// reverse_of_word = string(runes)
	// fmt.Println("given word : ", word)
	// fmt.Println("reverse of word : ", reverse_of_word)
	
	for _, letter := range word {
		reverse_of_word = string(letter) + reverse_of_word
	}
	fmt.Println("reverse of word : ", reverse_of_word)
}
