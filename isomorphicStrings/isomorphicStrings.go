package main

import "fmt"

func isIsomorphic(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    // maps for forward and reverse mapping
    sToT := make(map[byte]byte)
    tToS := make(map[byte]byte)

    for i := 0; i < len(s); i++ {
        c1, c2 := s[i], t[i]

        // check mapping from s -> t
        if mapped, ok := sToT[c1]; ok {
            if mapped != c2 {
                return false
            }
        } else {
            sToT[c1] = c2
        }

        // check mapping from t -> s
        if mapped, ok := tToS[c2]; ok {
            if mapped != c1 {
                return false
            }
        } else {
            tToS[c2] = c1
        }
    }

    return true
}

func main() {
    fmt.Println(isIsomorphic("egg", "add"))   // true
    fmt.Println(isIsomorphic("foo", "bar"))   // false
    fmt.Println(isIsomorphic("paper", "title")) // true
}
