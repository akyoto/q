package types

// Base is used to describe basic types like integers and floats.
type Base struct {
	name string
	size int
}

// Name returns the name of the type.
func (s *Base) Name() string {
	return s.name
}

// Size returns the total size in bytes.
func (s *Base) Size() int {
	return s.size
}