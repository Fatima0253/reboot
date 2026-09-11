package piscine

func ConcatSlice(slice1, slice2 []int) []int {
	var res []int

	for _, v := range slice1 {
		res = append(res, v)
	}

	for _, v := range slice2 {
		res = append(res, v)
	}

	return res
}