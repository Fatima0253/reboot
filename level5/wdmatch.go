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

	s1 := []rune(args[0])
	s2 := []rune(args[1])

	i := 0
	j := 0

	// Walk through s2 matching characters of s1 in order
	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			i++
		}
		j++
	}

	// If all characters of s1 were matched in order
	if i == len(s1) {
		for _, r := range s1 {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}