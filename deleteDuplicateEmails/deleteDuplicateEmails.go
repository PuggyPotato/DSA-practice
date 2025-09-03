package main

import (
	"fmt"
)

// Define a struct to represent a row in the Person table
type Person struct {
	id    int
	email string
}

func deleteDuplicateEmails(persons []Person) []Person {
	seen := make(map[string]int) // email -> smallest id
	result := []Person{}

	for _, p := range persons {
		if _, exists := seen[p.email]; !exists {
			// First time seeing this email
			seen[p.email] = p.id
			result = append(result, p)
		} else {
			// Already seen: keep the one with smaller id
			if p.id < seen[p.email] {
				// Replace in result
				for i := range result {
					if result[i].email == p.email {
						result[i] = p
						seen[p.email] = p.id
						break
					}
				}
			}
		}
	}
	return result
}

func main() {
	// Example test case
	persons := []Person{
		{1, "john@example.com"},
		{2, "bob@example.com"},
		{3, "john@example.com"},
	}

	result := deleteDuplicateEmails(persons)

	// Print result (order doesn’t matter)
	for _, p := range result {
		fmt.Printf("{%d, %s}\n", p.id, p.email)
	}
}
