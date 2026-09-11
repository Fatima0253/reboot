package piscine

func Slice(a []string, nbrs ...int) []string {
	if len(nbrs) == 0 {
		return nil
	}

	start := nbrs[0]
	end := len(a)

	if len(nbrs) > 1 {
		end = nbrs[1]
	}

	length := len(a)

	// Convert negative start index to positive equivalent
	if start < 0 {
		start += length
	}

	// Convert negative end index to positive equivalent
	if end < 0 {
		end += length
	}

	// Boundary adjustments
	if start < 0 {
		start = 0
	}

	if end > length {
		end = length
	}

	// Invalid slice range check
	if start >= end || start >= length || end < 0 {
		return nil
	}

	return a[start:end]
}