package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		return
	}

	s1 := args[0]
	s2 := args[1]

	// Track characters present in s2
	inS2 := make(map[rune]bool)
	for _, char := range s2 {
		inS2[char] = true
	}

	// Track characters already printed to avoid duplicates
	printed := make(map[rune]bool)

	for _, char := range s1 {
		if inS2[char] && !printed[char] {
			z01.PrintRune(char)
			printed[char] = true
		}
	}

	z01.PrintRune('\n')
}