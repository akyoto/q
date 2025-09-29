package data

// Data saves slices of bytes referenced by labels.
type Data struct {
	Immutable map[string][]byte
	Mutable   map[string][]byte
}