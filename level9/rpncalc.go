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

	if len(args) != 1 {
		printStr("Error\n")
		return
	}

	tokens := strings.Fields(args[0])
	if len(tokens) == 0 {
		printStr("Error\n")
		return
	}

	var stack []int

	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" || token == "%" {
			if len(stack) < 2 {
				printStr("Error\n")
				return
			}

			b := stack[len(stack)-1]
			a := stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			res := 0
			switch token {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				if b == 0 {
					printStr("Error\n")
					return
				}
				res = a / b
			case "%":
				if b == 0 {
					printStr("Error\n")
					return
				}
				res = a % b
			}
			stack = append(stack, res)
		} else {
			val, ok := atoi(token)
			if !ok {
				printStr("Error\n")
				return
			}
			stack = append(stack, val)
		}
	}

	if len(stack) != 1 {
		printStr("Error\n")
		return
	}

	printInt(stack[0])
	z01.PrintRune('\n')
}