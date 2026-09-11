package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		z01.PrintRune('\n')
		return
	}

	str := args[0]

	// Extract all non-empty words separated by spaces or tabs
	var words []string
	currentWord := ""

	for _, char := range str {
		if char == ' ' || char == '\t' {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(char)
		}
	}
	if currentWord != "" {
		words = append(words, currentWord)
	}

	if len(words) == 0 {
		z01.PrintRune('\n')
		return
	}

	// Rotate left: words from index 1 to end first, then index 0
	rotated := append(words[1:], words[0])

	// Print rotated words separated by a single space
	for i, word := range rotated {
		for _, r := range word {
			z01.PrintRune(r)
		}
		if i < len(rotated)-1 {
			z01.PrintRune(' ')
		}
	}

	z01.PrintRune('\n')
}