package main

import "fmt"

func missingNumber(nums []int) int {
    n := len(nums)
    expectedSum := n * (n + 1) / 2
    actualSum := 0
    for _, num := range nums {
        actualSum += num
    }
    return expectedSum - actualSum
}

func main() {
    nums1 := []int{3, 0, 1}
    fmt.Println(missingNumber(nums1)) // Output: 2

    nums2 := []int{0, 1}
    fmt.Println(missingNumber(nums2)) // Output: 2

    nums3 := []int{9, 6, 4, 2, 3, 5, 7, 0, 1}
    fmt.Println(missingNumber(nums3)) // Output: 8
}
