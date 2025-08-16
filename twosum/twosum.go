package main

import "fmt"

// twoSum finds indices of two numbers in nums that add up to target
func twoSum(nums []int, target int) []int {
    m := make(map[int]int) // map to store value -> index
    for i, num := range nums {
        complement := target - num
        if idx, found := m[complement]; found {
            return []int{idx, i}
        }
        m[num] = i
    }
    return nil
}

func main() {
    nums := []int{2, 7, 11, 15}
    target := 9
    result := twoSum(nums, target)
    fmt.Println(result) // Output: [0 1]
}
