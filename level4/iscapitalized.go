package piscine

func IsCapitalized(s string) bool {
	if len(s) == 0 {
		return false
	}

	inWord := false

	for _, char := range s {
		if char != ' ' && char != '\t' {
			if !inWord {
				// We are at the first character of a word
				if char >= 'a' && char <= 'z' {
					return false
				}
				inWord = true
			}
		} else {
			inWord = false
		}
	}

	return true
}