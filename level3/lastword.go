package piscine

func LastWord(s string) string {
	end := -1
	start := -1

	// Scan backwards from the end of the string
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			if end == -1 {
				end = i + 1
			}
		} else if end != -1 {
			start = i + 1
			break
		}
	}

	if end == -1 {
		return "\n"
	}

	if start == -1 {
		start = 0
	}

	return s[start:end] + "\n"
}