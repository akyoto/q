package data

// Insert registers a slice of bytes for the given label.
func (data Data) Insert(label string, raw []byte) {
	data[label] = raw
}