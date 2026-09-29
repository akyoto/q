package data

// appendMutable adds data that is mutable and not subject to string interning.
func (data *Data) appendMutable(final []byte, positions map[string]int) []byte {
	return appendLabeled(final, positions, data.Mutable)
}