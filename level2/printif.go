package piscine

func PrintIf(str string) string {
	if str == "" || len(str) >= 3 {
		return "G\n"
	}
	return "Invalid Input\n"
}