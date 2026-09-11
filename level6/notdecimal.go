package piscine

func NotDecimal(dec string) string {
	if dec == "" {
		return "\n"
	}

	// Find the position of the decimal point
	dotIdx := -1
	for i, char := range dec {
		if char == '.' {
			dotIdx = i
			break
		}
	}

	// No decimal point found
	if dotIdx == -1 {
		return dec + "\n"
	}

	// Check if string contains invalid characters (must be digits, optional leading '-', or single '.')
	dotCount := 0
	for i, char := range dec {
		if char == '-' && i == 0 {
			continue
		}
		if char == '.' {
			dotCount++
			if dotCount > 1 {
				return dec + "\n"
			}
			continue
		}
		if char < '0' || char > '9' {
			return dec + "\n"
		}
	}

	afterDot := dec[dotIdx+1:]

	// If nothing after dot, or only '0' after dot (e.g. "1952.0" or "0.")
	if afterDot == "" || afterDot == "0" {
		return dec + "\n"
	}

	// Construct result without the '.'
	res := dec[:dotIdx] + afterDot

	// Strip leading zeros while preserving sign
	sign := ""
	if len(res) > 0 && res[0] == '-' {
		sign = "-"
		res = res[1:]
	}

	// Trim leading zeros
	start := 0
	for start < len(res)-1 && res[start] == '0' {
		start++
	}
	res = res[start:]

	return sign + res + "\n"
}