package piscine

func CanJump(a []uint) bool {
	if len(a) == 0 {
		return false
	}

	currentIndex := 0
	lastIndex := len(a) - 1

	for currentIndex < lastIndex {
		steps := int(a[currentIndex])

		// If steps is 0, we are stuck and cannot advance
		if steps == 0 {
			return false
		}

		currentIndex += steps
	}

	// We return true only if we landed exactly on the last index
	return currentIndex == lastIndex
}