package piscine

func RepeatAlpha(s string) string {
	var res []rune

	for _, r := range s {
		count := 1

		if r >= 'a' && r <= 'z' {
			count = int(r - 'a' + 1)
		} else if r >= 'A' && r <= 'Z' {
			count = int(r - 'A' + 1)
		}

		for i := 0; i < count; i++ {
			res = append(res, r)
		}
	}

	return string(res)
}