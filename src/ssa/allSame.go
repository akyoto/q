package ssa

// allSame checks if all elements are the same.
func allSame[T comparable](slice []T) bool {
	if len(slice) <= 1 {
		return true
	}

	first := slice[0]

	for _, v := range slice[1:] {
		if v != first {
			return false
		}
	}

	return true
}