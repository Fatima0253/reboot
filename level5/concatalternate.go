package piscine

func ConcatAlternate(slice1, slice2 []int) []int {
	first := slice1
	second := slice2

	// Determine which slice goes first based on length criteria
	if len(slice2) > len(slice1) {
		first = slice2
		second = slice1
	}

	var res []int
	i, j := 0, 0

	// Alternate elements as long as both slices have remaining elements
	for i < len(first) && j < len(second) {
		res = append(res, first[i])
		res = append(res, second[j])
		i++
		j++
	}

	// Append any remaining elements from the longer slice
	for i < len(first) {
		res = append(res, first[i])
		i++
	}
	for j < len(second) {
		res = append(res, second[j])
		j++
	}

	return res
}