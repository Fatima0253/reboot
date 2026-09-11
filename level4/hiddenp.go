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

	// An empty string s1 is always hidden in s2
	if len(s1) == 0 {
		z01.PrintRune('1')
		z01.PrintRune('\n')
		return
	}

	i := 0 // Index pointer for s1
	s1Runes := []rune(s1)

	for _, char := range s2 {
		if char == s1Runes[i] {
			i++
			if i == len(s1Runes) {
				z01.PrintRune('1')
				z01.PrintRune('\n')
				return
			}
		}
	}

	z01.PrintRune('0')
	z01.PrintRune('\n')
}