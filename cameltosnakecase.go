package piscine

func CamelToSnakeCase(s string) string {
	if s == "" {
		return ""
	}

	// 1. Validate if string follows camelCase rules
	runes := []rune(s)
	n := len(runes)

	for i := 0; i < n; i++ {
		r := runes[i]

		// Numbers or punctuation are not allowed
		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
			return s
		}

		// Capitalized letter rules
		if r >= 'A' && r <= 'Z' {
			// Word cannot end on a capitalized letter
			if i == n-1 {
				return s
			}
			// No two capitalized letters directly adjacent
			if i+1 < n && (runes[i+1] >= 'A' && runes[i+1] <= 'Z') {
				return s
			}
		}
	}

	// 2. Convert to snake_case
	var res []rune
	for i := 0; i < n; i++ {
		r := runes[i]
		if r >= 'A' && r <= 'Z' {
			if i != 0 {
				res = append(res, '_')
			}
			res = append(res, r)
		} else {
			res = append(res, r)
		}
	}

	return string(res)
}