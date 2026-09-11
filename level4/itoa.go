package piscine

func Itoa(n int) string {
	if n == 0 {
		return "0"
	}

	isNegative := false
	if n < 0 {
		isNegative = true
		n = -n
	}

	var runes []rune
	for n > 0 {
		digit := n % 10
		runes = append([]rune{rune('0' + digit)}, runes...)
		n /= 10
	}

	if isNegative {
		runes = append([]rune{'-'}, runes...)
	}

	return string(runes)
}