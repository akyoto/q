package types

// Type represents a type in the type system.
type Type struct {
	Name string
	Size int
}

// String returns the type name.
func (typ *Type) String() string {
	if typ == nil {
		return "unknown type"
	}

	return typ.Name
}
