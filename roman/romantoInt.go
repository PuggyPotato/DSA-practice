package main

import (
    "fmt"
)

func romanToInt(s string) int {
    // Map of Roman numerals to integers
    romanMap := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    result := 0
    n := len(s)

    for i := 0; i < n; i++ {
        // If this numeral is smaller than the next, subtract it
        if i < n-1 && romanMap[s[i]] < romanMap[s[i+1]] {
            result -= romanMap[s[i]]
        } else {
            result += romanMap[s[i]]
        }
    }

    return result
}

func main() {
    fmt.Println(romanToInt("III"))    // 3
    fmt.Println(romanToInt("IV"))     // 4
    fmt.Println(romanToInt("IX"))     // 9
    fmt.Println(romanToInt("LVIII"))  // 58
    fmt.Println(romanToInt("MCMXCIV"))// 1994
}
