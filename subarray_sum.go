// You can edit this code!
// Click here and start typing.
// find largest subarray sum
package main

import "fmt"

func maxSubarraySum(numbers []int) int{
	current_sum, max_sum := 0, 0
	for _, num := range numbers {
		fmt.Println(num)
		current_sum = max(num, current_sum+num)
		max_sum = max(max_sum, current_sum)

	}
	return max_sum
}
func main() {
	fmt.Println("Largest subarray sum")
	num_list := []int{11, 2, 5, -11, 0, 1}
	fmt.Println("max_sum :", maxSubarraySum(num_list))
}
