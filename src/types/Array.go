package types

// Array is the address of an object.
type Array struct {
	Of Type
}

// Name returns the type name.
func (a *Array) Name() string {
	return "[]" + a.Of.Name()
}

// Size returns the total size in bytes.
func (a *Array) Size() int {
	return 8
}