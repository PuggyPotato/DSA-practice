package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	n := len(digits)

	// Traverse from last digit to first
	for i := n - 1; i >= 0; i-- {
		if digits[i] < 9 {
			digits[i]++
			return digits
		}
		// If digit is 9, set to 0 and continue loop
		digits[i] = 0
	}

	// If all digits were 9 (like 999 -> 1000), create new slice
	result := make([]int, n+1)
	result[0] = 1
	return result
}

func main() {
	fmt.Println(plusOne([]int{1, 2, 3})) // [1 2 4]
	fmt.Println(plusOne([]int{4, 3, 2, 1})) // [4 3 2 2]
	fmt.Println(plusOne([]int{9})) // [1 0]
	fmt.Println(plusOne([]int{9, 9, 9})) // [1 0 0 0]
}
