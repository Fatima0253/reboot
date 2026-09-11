package piscine

func RevConcatAlternate(slice1, slice2 []int) []int {
	var res []int

	// Standardize which slice is longer and which is shorter
	first := slice1
	second := slice2

	if len(slice2) > len(slice1) {
		first = slice2
		second = slice1
	}

	i := len(first) - 1
	j := len(second) - 1

	// Step 1: Append elements from the longer slice until both remaining parts are equal in length
	for i > j {
		res = append(res, first[i])
		i--
	}

	// Step 2: Alternate elements from slice1 first, then slice2
	// If slice1 was originally shorter/equal, 'first' is slice1 or slice2 respectively.
	// Always alternate taking from slice1's remaining end, then slice2's remaining end.
	i1 := i
	j1 := j

	if len(slice2) > len(slice1) {
		// When slice2 was longer, first=slice2 (index i1) and second=slice1 (index j1)
		for j1 >= 0 {
			res = append(res, second[j1]) // slice1 element
			res = append(res, first[i1])  // slice2 element
			j1--
			i1--
		}
	} else {
		// When slice1 was longer or equal, first=slice1 (index i1) and second=slice2 (index j1)
		for i1 >= 0 {
			res = append(res, first[i1]) // slice1 element
			if j1 >= 0 {
				res = append(res, second[j1]) // slice2 element
				j1--
			}
			i1--
		}
	}

	return res
}