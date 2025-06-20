package types

// Pointer is the address of an object.
type Pointer struct {
	To Type
}

// Name returns the type name.
func (p *Pointer) Name() string {
	return "*" + p.To.Name()
}

// Size returns the total size in bytes.
func (p *Pointer) Size() int {
	return 8
}