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
	if str == "" {
		z01.PrintRune('\n')
		return
	}

	// Split words by space
	var words []string
	currentWord := ""

	for _, char := range str {
		if char == ' ' {
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

	// Print words in reverse order with spaces in between
	for i := len(words) - 1; i >= 0; i-- {
		for _, char := range words[i] {
			z01.PrintRune(char)
		}
		if i > 0 {
			z01.PrintRune(' ')
		}
	}

	z01.PrintRune('\n')
}