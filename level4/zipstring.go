package piscine

func ZipString(s string) string {
	if len(s) == 0 {
		return ""
	}

	runes := []rune(s)
	var res string

	i := 0
	for i < len(runes) {
		currentChar := runes[i]
		count := 0

		// Count consecutive occurrences of the current character
		for i < len(runes) && runes[i] == currentChar {
			count++
			i++
		}

		// Convert count to string digits
		countStr := ""
		temp := count
		for temp > 0 {
			countStr = string(rune('0'+temp%10)) + countStr
			temp /= 10
		}

		res += countStr + string(currentChar)
	}

	return res
}