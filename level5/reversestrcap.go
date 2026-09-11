package main

import (
	"os"

	"github.com/01-edu/z01"
)

func processWord(s string) string {
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n; i++ {
		char := runes[i]

		// First, convert all letters to lowercase
		if char >= 'A' && char <= 'Z' {
			char = char + ('a' - 'A')
		}

		// Check if it's an alphabetic character
		if char >= 'a' && char <= 'z' {
			// Capitalize if it's the last character of string OR followed by a space/tab
			isLastOfWord := (i == n-1) || (runes[i+1] == ' ' || runes[i+1] == '\t')
			if isLastOfWord {
				char = char - ('a' - 'A')
			}
		}

		runes[i] = char
	}

	return string(runes)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		return
	}

	for _, arg := range args {
		result := processWord(arg)
		for _, r := range result {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}