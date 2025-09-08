package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	// Compile regex for both formats
	pattern := regexp.MustCompile(`^(\(\d{3}\)\s\d{3}-\d{4}|\d{3}-\d{3}-\d{4})$`)

	// Open the file
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Scan each line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if pattern.MatchString(line) {
			fmt.Println(line)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}
}
