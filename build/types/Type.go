package types

// Type represents a type in the type system.
type Type struct {
	Name   string
	Size   uint
	Fields []*Field
}

// FieldByName returns the field with the given name.
func (typ *Type) FieldByName(name string) *Field {
	for _, field := range typ.Fields {
		if field.Name == name {
			return field
		}
	}

	return nil
}

// String returns the type name.
func (typ *Type) String() string {
	if typ == nil {
		return "unknown type"
	}

	return typ.Name
}
