package main

import "fmt"

// removeDuplicates removes duplicates in-place and returns the new length.
func removeDuplicates(nums []int) int {
    if len(nums) == 0 {
        return 0
    }

    // slow pointer - the index of the last unique element
    slow := 0

    // fast pointer - scans through the array
    for fast := 1; fast < len(nums); fast++ {
        if nums[fast] != nums[slow] {
            slow++
            nums[slow] = nums[fast]
        }
    }

    return slow + 1
}

func main() {
    nums := []int{0,0,1,1,1,2,2,3,3,4}
    length := removeDuplicates(nums)

    fmt.Println("New length:", length)
    fmt.Println("Array without duplicates:", nums[:length])
}
