package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	// Negative numbers are not palindrome
	if x < 0 {
		return false
	}

	original := x
	reversed := 0

	for x != 0 {
		digit := x % 10
		reversed = reversed*10 + digit
		x /= 10
	}

	return original == reversed
}

func main() {
	num := 121
	fmt.Println(isPalindrome(num)) // true

	num = -121
	fmt.Println(isPalindrome(num)) // false

	num = 10
	fmt.Println(isPalindrome(num)) // false
}
