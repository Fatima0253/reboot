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

func main() {
	args := os.Args[1:]

	// 1. No arguments provided
	if len(args) == 0 {
		printStr("options: abcdefghijklmnopqrstuvwxyz\n")
		return
	}

	// 2. Check for -h priority on the FIRST flag containing 'h'
	// If any argument starts with '-' and the FIRST character after '-' is 'h', print usage.
	for _, arg := range args {
		if len(arg) > 1 && arg[0] == '-' {
			if arg[1] == 'h' {
				printStr("options: abcdefghijklmnopqrstuvwxyz\n")
				return
			}
			break
		}
	}

	var flags uint32 = 0

	// 3. Process all flags and validate option characters
	for _, arg := range args {
		// Single '-' or string not starting with '-' is invalid
		if len(arg) < 2 || arg[0] != '-' {
			printStr("Invalid Option\n")
			return
		}

		for i := 1; i < len(arg); i++ {
			c := arg[i]
			if c >= 'a' && c <= 'z' {
				// Map 'a' -> bit 0, 'b' -> bit 1, ..., 'z' -> bit 25
				flags |= (1 << (c - 'a'))
			} else {
				// Character other than a-z is an invalid option
				printStr("Invalid Option\n")
				return
			}
		}
	}

	// 4. Output bitwise flags as 4 groups of 8 binary digits
	for i := 31; i >= 0; i-- {
		if (flags & (1 << i)) != 0 {
			z01.PrintRune('1')
		} else {
			z01.PrintRune('0')
		}

		if i%8 == 0 && i != 0 {
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune('\n')
}