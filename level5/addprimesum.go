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

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
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
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	num, ok := atoi(args[0])
	if !ok || num <= 0 {
		z01.PrintRune('0')
		z01.PrintRune('\n')
		return
	}

	sum := 0
	for i := 2; i <= num; i++ {
		if isPrime(i) {
			sum += i
		}
	}

	printNum(sum)
	z01.PrintRune('\n')
}