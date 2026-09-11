package piscine

func WordFlip(str string) string {
	if str == "" {
		return "Invalid Output\n"
	}

	// Extract non-empty words while ignoring multiple spaces/tabs
	var words []string
	currentWord := ""

	for _, char := range str {
		if char == ' ' || char == '\t' {
			if currentWord != "" {
				words = append(words, currentWord)
				currentWord = ""
			}
		} else {
			currentWord += string(char)
		}
	}
	if currentWord != "" {
		words = append(words, currentWord)
	}

	// If there were only spaces/tabs in the input string
	if len(words) == 0 {
		return "Invalid Output\n"
	}

	// Build the reversed string
	res := ""
	for i := len(words) - 1; i >= 0; i-- {
		res += words[i]
		if i > 0 {
			res += " "
		}
	}

	return res + "\n"
}