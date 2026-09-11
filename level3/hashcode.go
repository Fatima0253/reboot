package piscine

func HashCode(dec string) string {
	size := len(dec)
	var res []rune

	for _, char := range dec {
		hashed := (int(char) + size) % 127
		if hashed < 32 || hashed == 127 {
			hashed += 33
		}
		res = append(res, rune(hashed))
	}

	return string(res)
}