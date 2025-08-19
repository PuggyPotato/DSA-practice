package main

import (
    "fmt"
)

func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    prefix := strs[0] // start with the first string as prefix
    for _, str := range strs[1:] {
        // Reduce prefix until it matches the start of str
        for len(prefix) > 0 && str[:len(prefix)] != prefix {
            prefix = prefix[:len(prefix)-1] // chop off last character
        }
        if prefix == "" {
            return ""
        }
    }
    return prefix
}

func main() {
    strs := []string{"flower", "flow", "flight"}
    fmt.Println("Longest Common Prefix:", longestCommonPrefix(strs)) // Output: "fl"
}
