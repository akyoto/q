package types

// Type is the generic interface for different data types.
type Type interface {
	Name() string
	Size() int
}