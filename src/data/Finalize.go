package data

// Finalize returns the final raw data slice and a map of labels with their respective indices.
// It will try to reuse existing data whenever possible.
func (data *Data) Finalize() ([]byte, map[string]int) {
	capacity := 0

	for _, value := range data.Immutable {
		capacity += len(value)
	}

	for _, value := range data.Mutable {
		capacity += len(value)
	}

	final := make([]byte, 0, capacity)
	positions := make(map[string]int, len(data.Immutable)+len(data.Mutable))
	final = data.appendImmutable(final, positions)
	final = data.appendMutable(final, positions)
	return final, positions
}