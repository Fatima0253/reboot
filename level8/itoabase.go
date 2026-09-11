package piscine

func ItoaBase(value, base int) string {
	if base < 2 || base > 16 {
		return ""
	}

	if value == 0 {
		return "0"
	}

	baseChars := "0123456789ABCDEF"
	isNegative := false

	if value < 0 {
		isNegative = true
	}

	var res []byte

	// Process digits using unsigned/absolute value handling
	for value != 0 {
		remainder := value % base
		if remainder < 0 {
			remainder = -remainder
		}
		res = append([]byte{baseChars[remainder]}, res...)
		value /= base
	}

	if isNegative {
		res = append([]byte{'-'}, res...)
	}

	return string(res)
}