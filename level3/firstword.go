package piscine

func FirstWord(s string) string {
	start := -1
	end := len(s)

	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			if start == -1 {
				start = i
			}
		} else if start != -1 {
			end = i
			break
		}
	}

	if start == -1 {
		return "\n"
	}

	return s[start:end] + "\n"
}