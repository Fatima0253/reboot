package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
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

	// If there are no words, display nothing
	if len(words) == 0 {
		return
	}

	// Print words with exactly three spaces between each word
	for i, word := range words {
		if i > 0 {
			z01.PrintRune(' ')
			z01.PrintRune(' ')
			z01.PrintRune(' ')
		}
		for _, char := range word {
			z01.PrintRune(char)
		}
	}
	z01.PrintRune('\n')
}