package main

import (
	"os"

	"github.com/01-edu/z01"
)

func atoi(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	n := 0
	for _, char := range s {
		if char < '0' || char > '9' {
			return 0, false
		}
		n = n*10 + int(char-'0')
	}
	return n, true
}

func printNum(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	var digits []rune
	for n > 0 {
		digits = append([]rune{rune('0' + n%10)}, digits...)
		n /= 10
	}
	for _, d := range digits {
		z01.PrintRune(d)
	}
}

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		return
	}

	num, ok := atoi(args[0])
	if !ok || num <= 1 {
		return
	}

	first := true
	factor := 2

	for num > 1 {
		if num%factor == 0 {
			if !first {
				z01.PrintRune('*')
			}
			printNum(factor)
			first = false
			num /= factor
		} else {
			factor++
		}
	}

	z01.PrintRune('\n')
}