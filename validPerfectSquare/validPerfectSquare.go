package main

import (
	"fmt"
	"strings"
)

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	if len(pattern) != len(words) {
		return false
	}

	// Maps for letter -> word and word -> letter
	letterToWord := make(map[byte]string)
	wordToLetter := make(map[string]byte)

	for i := 0; i < len(pattern); i++ {
		letter := pattern[i]
		word := words[i]

		// Check letter -> word mapping
		if mappedWord, ok := letterToWord[letter]; ok {
			if mappedWord != word {
				return false
			}
		} else {
			letterToWord[letter] = word
		}

		// Check word -> letter mapping
		if mappedLetter, ok := wordToLetter[word]; ok {
			if mappedLetter != letter {
				return false
			}
		} else {
			wordToLetter[word] = letter
		}
	}

	return true
}

func main() {
	fmt.Println(wordPattern("abba", "dog cat cat dog")) // true
	fmt.Println(wordPattern("abba", "dog cat cat fish")) // false
	fmt.Println(wordPattern("aaaa", "dog cat cat dog")) // false
	fmt.Println(wordPattern("abc", "dog cat dog")) // false
}
