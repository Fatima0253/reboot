package piscine

func WeAreUnique(str1, str2 string) int {
	if str1 == "" && str2 == "" {
		return -1
	}

	seen1 := make(map[rune]bool)
	seen2 := make(map[rune]bool)

	for _, char := range str1 {
		seen1[char] = true
	}

	for _, char := range str2 {
		seen2[char] = true
	}

	count := 0

	// Count characters in str1 that are NOT in str2
	for char := range seen1 {
		if !seen2[char] {
			count++
		}
	}

	// Count characters in str2 that are NOT in str1
	for char := range seen2 {
		if !seen1[char] {
			count++
		}
	}

	return count
}