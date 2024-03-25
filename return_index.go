// You can edit this code!
// Click here and start typing.
// find target sum number indices
package main

import "fmt"
func main() {
	n := []int{1, 2, 3}
	fmt.Println(n)
	target := 4
	var compliemnt_map map[int]int
	compliemnt_map = map[int]int{} // initialization
	// compliemnt_map = make(map[int]int)  // type of initialization create empty map
	for ind, num := range n {
		compliment := target - num
		if idx, exists := compliemnt_map[compliment]; exists {
			fmt.Println("found", ind, idx)
		}
		compliemnt_map[num] = ind
	}

}
