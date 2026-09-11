package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		z01.PrintRune('\n')
		return
	}

	seen := make(map[rune]bool)

	// Process first string
	for _, char := range args[0] {
		if !seen[char] {
			z01.PrintRune(char)
			seen[char] = true
		}
	}

	// Process second string
	for _, char := range args[1] {
		if !seen[char] {
			z01.PrintRune(char)
			seen[char] = true
		}
	}

	z01.PrintRune('\n')
}