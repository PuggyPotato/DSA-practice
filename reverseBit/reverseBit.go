package main

import "fmt"

// reverseBits reverses the bits of a 32-bit unsigned integer
func reverseBits(n uint32) uint32 {
    var result uint32 = 0
    for i := 0; i < 32; i++ {
        result <<= 1           // make room for the next bit
        result |= n & 1        // add the last bit of n
        n >>= 1                // shift n to process the next bit
    }
    return result
}

func main() {
    n1 := uint32(43261596)
    n2 := uint32(2147483644)

    fmt.Println(reverseBits(n1)) // Output: 964176192
    fmt.Println(reverseBits(n2)) // Output: 1073741822
}
