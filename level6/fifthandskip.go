package piscine

func FifthAndSkip(str string) string {
	if str == "" {
		return "\n"
	}

	// Filter out spaces first
	var filtered []rune
	for _, char := range str {
		if char != ' ' {
			filtered = append(filtered, char)
		}
	}

	if len(filtered) < 5 {
		return "Invalid Input\n"
	}

	var res []rune
	count := 0

	for i := 0; i < len(filtered); i++ {
		count++
		if count == 6 {
			// Skip the 6th character and reset count
			count = 0
			continue
		}

		res = append(res, filtered[i])

		// Add space after every group of 5 characters (unless it's the end)
		if count == 5 && i < len(filtered)-1 {
			res = append(res, ' ')
		}
	}

	return string(res) + "\n"
}