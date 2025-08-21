package main

import "fmt"

func removeElement(nums []int, val int) int {
    k := 0 // pointer for placing non-val elements
    for _, num := range nums {
        if num != val {
            nums[k] = num
            k++
        }
    }
    return k
}

func main() {
    nums1 := []int{3, 2, 2, 3}
    val1 := 3
    k1 := removeElement(nums1, val1)
    fmt.Println("New length:", k1)
    fmt.Println("Modified array:", nums1[:k1])

    nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
    val2 := 2
    k2 := removeElement(nums2, val2)
    fmt.Println("New length:", k2)
    fmt.Println("Modified array:", nums2[:k2])
}
