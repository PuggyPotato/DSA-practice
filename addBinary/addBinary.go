package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	i, j := len(a)-1, len(b)-1
	carry := 0
	result := []byte{}

	for i >= 0 || j >= 0 || carry > 0 {
		sum := carry
		if i >= 0 {
			sum += int(a[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(b[j] - '0')
			j--
		}
		result = append([]byte{byte(sum%2 + '0')}, result...)
		carry = sum / 2
	}

	return string(result)
}

func main() {
	a := "1010"
	b := "1011"
	fmt.Println(addBinary(a, b)) // Output: 10101
}
