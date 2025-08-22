package main

import (
	"fmt"
)

func mySqrt(x int) int {
	if x < 2 {
		return x
	}

	left, right := 1, x/2
	var ans int

	for left <= right {
		mid := left + (right-left)/2
		if mid <= x/mid { // Avoid overflow by using division instead of mid*mid
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return ans
}

func main() {
	fmt.Println(mySqrt(4)) // Output: 2
	fmt.Println(mySqrt(8)) // Output: 2
}
