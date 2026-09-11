package piscine

func FromTo(from int, to int) string {
	if from < 0 || from > 99 || to < 0 || to > 99 {
		return "Invalid\n"
	}

	formatNum := func(n int) string {
		if n < 10 {
			return string(rune('0')) + string(rune('0'+n))
		}
		return string(rune('0'+n/10)) + string(rune('0'+n%10))
	}

	var res string
	step := 1
	if from > to {
		step = -1
	}

	i := from
	for {
		res += formatNum(i)
		if i == to {
			break
		}
		res += ", "
		i += step
	}

	return res + "\n"
}