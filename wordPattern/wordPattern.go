package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(words) != len(pattern) {
		return false
	}

	// Maps for bijection
	charToWord := make(map[byte]string)
	wordToChar := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		ch := pattern[i]
		word := words[i]

		// If the char already maps to a word, check consistency
		if w, ok := charToWord[ch]; ok {
			if w != word {
				return false
			}
		} else {
			charToWord[ch] = word
		}

		// If the word already maps to a char, check consistency
		if c, ok := wordToChar[word]; ok {
			if c != ch {
				return false
			}
		} else {
			wordToChar[word] = ch
		}
	}

	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog"))   // true
	fmt.Println(wordPattern("abba", "dog cat cat fish"))  // false
	fmt.Println(wordPattern("aaaa", "dog cat cat dog"))   // false
	fmt.Println(wordPattern("abba", "dog dog dog dog"))   // false
}
