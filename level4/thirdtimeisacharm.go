package piscine

func ThirdTimeIsACharm(str string) string {
	if len(str) < 3 {
		return "\n"
	}

	var res []rune
	runes := []rune(str)

	for i := 2; i < len(runes); i += 3 {
		res = append(res, runes[i])
	}

	return string(res) + "\n"
}