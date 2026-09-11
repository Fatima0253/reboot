package main

import (
	"os"
	"strings"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func printInt(n int) {
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	if n < 0 {
		z01.PrintRune('-')
		n = -n
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

func atoi(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	sign, i := 1, 0
	if s[0] == '-' {
		sign, i = -1, 1
	} else if s[0] == '+' {
		i = 1
	}
	if i >= len(s) {
		return 0, false
	}
	n := 0
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		n = n*10 + int(s[i]-'0')
	}
	return sign * n, true
}

func main() {
	args := os.Args[1:]

	if len(args) != 2 {
		printStr("Invalid input.\n")
		return
	}

	target, ok := atoi(args[1])
	if !ok {
		printStr("Invalid target sum.\n")
		return
	}

	arrStr := args[0]
	if !strings.HasPrefix(arrStr, "[") || !strings.HasSuffix(arrStr, "]") {
		printStr("Invalid input.\n")
		return
	}

	// Remove '[' and ']'
	content := arrStr[1 : len(arrStr)-1]
	if content == "" {
		printStr("No pairs found.\n")
		return
	}

	parts := strings.Split(content, ",")
	var arr []int

	for _, p := range parts {
		token := strings.TrimSpace(p)
		val, ok := atoi(token)
		if !ok {
			printStr("Invalid number: " + token + "\n")
			return
		}
		arr = append(arr, val)
	}

	// Find pair indices
	type pair struct{ i, j int }
	var pairs []pair

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]+arr[j] == target {
				pairs = append(pairs, pair{i, j})
			}
		}
	}

	if len(pairs) == 0 {
		printStr("No pairs found.\n")
		return
	}

	printStr("Pairs with sum ")
	printInt(target)
	printStr(": [")

	for idx, p := range pairs {
		if idx > 0 {
			z01.PrintRune(' ')
		}
		z01.PrintRune('[')
		printInt(p.i)
		z01.PrintRune(' ')
		printInt(p.j)
		z01.PrintRune(']')
	}
	printStr("]\n")
}