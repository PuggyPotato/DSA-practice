package main

import "fmt"

func isPowerOfTwo(n int) bool {
    return n > 0 && (n & (n - 1)) == 0
}

func main() {
    fmt.Println(isPowerOfTwo(1))  // true
    fmt.Println(isPowerOfTwo(16)) // true
    fmt.Println(isPowerOfTwo(3))  // false
}
