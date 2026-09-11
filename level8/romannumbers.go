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

func atoi(s string) (int, bool) {
	if len(s) == 0 {
		return 0, false
	}
	n := 0
	for _, r := range s {
		if r < '0' || r > '9' {
			return 0, false
		}
		n = n*10 + int(r-'0')
	}
	return n, true
}

type romanMapping struct {
	val        int
	calculation string
	symbol     string
}

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		return
	}

	num, ok := atoi(args[0])
	if !ok || num <= 0 || num >= 4000 {
		printStr("ERROR: cannot convert to roman digit\n")
		return
	}

	mappings := []romanMapping{
		{1000, "M", "M"},
		{900, "(M-C)", "CM"},
		{500, "D", "D"},
		{400, "(D-C)", "CD"},
		{100, "C", "C"},
		{90, "(C-X)", "XC"},
		{50, "L", "L"},
		{40, "(L-X)", "XL"},
		{10, "X", "X"},
		{9, "(X-I)", "IX"},
		{5, "V", "V"},
		{4, "(V-I)", "IV"},
		{1, "I", "I"},
	}

	var calcParts []string
	var romanResult string

	temp := num
	for _, m := range mappings {
		for temp >= m.val {
			calcParts = append(calcParts, m.calculation)
			romanResult += m.symbol
			temp -= m.val
		}
	}

	// Print calculation parts joined by '+'
	for i, part := range calcParts {
		if i > 0 {
			z01.PrintRune('+')
		}
		printStr(part)
	}
	z01.PrintRune('\n')

	// Print final Roman numeral string
	printStr(romanResult + "\n")
}