package main

func is_palindrome(word string) bool{
	word_rune := []rune(word)
	word_len := len(word_rune)

	for i:=0;i<word_len/2;i++{
		if word_rune[i] != word_rune[word_len-1-i]{
			return false
		}
	}
	return true
}

func main(){
	word := "malayalam"
	if is_palindrome(word){
		fmt.Println("Number is palindrome")
	}else{
		fmt.Println("Number is not palindrome")
	}
}