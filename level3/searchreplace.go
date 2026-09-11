package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	if len(args) != 3 {
		return
	}

	str := args[0]
	target := args[1]
	replacement := args[2]

	// Handle case where target or replacement is not a single character
	if len(target) != 1 || len(replacement) != 1 {
		return
	}

	oldChar := rune(target[0])
	newChar := rune(replacement[0])

	for _, r := range str {
		if r == oldChar {
			z01.PrintRune(newChar)
		} else {
			z01.PrintRune(r)
		}
	}
	z01.PrintRune('\n')
}