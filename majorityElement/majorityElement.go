package main

import "fmt"

// majorityElement finds the majority element using Boyer-Moore Voting Algorithm
func majorityElement(nums []int) int {
    candidate := nums[0]
    count := 0

    for _, num := range nums {
        if count == 0 {
            candidate = num
        }
        if num == candidate {
            count++
        } else {
            count--
        }
    }
    return candidate
}

func main() {
    fmt.Println(majorityElement([]int{3, 2, 3}))        // Output: 3
    fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2})) // Output: 2
}
