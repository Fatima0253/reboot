package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func isVowel(r rune) bool {
	return r == 'a' || r == 'e' || r == 'i' || r == 'o' || r == 'u' ||
		r == 'A' || r == 'E' || r == 'I' || r == 'O' || r == 'U'
}

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		return
	}

	str := args[0]

	// Find the index of the first vowel
	firstVowelIdx := -1
	for i, char := range str {
		if isVowel(char) {
			firstVowelIdx = i
			break
		}
	}

	// No vowels found
	if firstVowelIdx == -1 {
		printStr("No vowels\n")
		return
	}

	// Transform word: [first_vowel_to_end] + [leading_consonants] + "ay"
	res := str[firstVowelIdx:] + str[:firstVowelIdx] + "ay"
	printStr(res + "\n")
}