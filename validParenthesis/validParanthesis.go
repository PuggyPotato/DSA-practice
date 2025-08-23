package main

import "fmt"

func isValid(s string) bool {
    stack := []rune{} // stack to store opening brackets

    // mapping of closing to opening brackets
    pairs := map[rune]rune{
        ')': '(',
        ']': '[',
        '}': '{',
    }

    for _, char := range s {
        switch char {
        case '(', '[', '{':
            // push opening brackets onto stack
            stack = append(stack, char)
        case ')', ']', '}':
            // if stack is empty or top of stack doesn't match
            if len(stack) == 0 || stack[len(stack)-1] != pairs[char] {
                return false
            }
            // pop the top element
            stack = stack[:len(stack)-1]
        }
    }

    // valid if stack is empty at the end
    return len(stack) == 0
}

func main() {
    tests := []string{"()", "()[]{}", "(]", "([])", "([)]"}
    for _, test := range tests {
        fmt.Printf("Input: %s -> Output: %v\n", test, isValid(test))
    }
}
