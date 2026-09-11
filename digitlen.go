package piscine

func DigitLen(n, base int) int {
	if base < 2 || base > 36 {
		return -1
	}

	if n < 0 {
		n = -n
	}

	count := 0
	for {
		n /= base
		count++
		if n == 0 {
			break
		}
	}

	return count
}