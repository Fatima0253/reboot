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
	var words []string
	var currentWord []rune

	// Parse string into words, delimited by spaces or tabs ('\t')
	for _, char := range str {
		if char == ' ' || char == '\t' {
			if len(currentWord) > 0 {
				words = append(words, string(currentWord))
				currentWord = nil
			}
		} else {
			currentWord = append(currentWord, char)
		}
	}

	if len(currentWord) > 0 {
		words = append(words, string(currentWord))
	}

	// If no words were found, output a single newline
	if len(words) == 0 {
		z01.PrintRune('\n')
		return
	}

	// Print words joined by a single space
	for i, word := range words {
		if i > 0 {
			z01.PrintRune(' ')
		}
		for _, char := range word {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}