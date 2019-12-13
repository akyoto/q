package types

import "github.com/akyoto/q/build/token"

// Field is a field in a data structure.
type Field struct {
	Name     string
	Type     *Type
	Position token.Position
	Offset   uint
	Mutable  bool
}

// String returns the name of the parameter.
func (field *Field) String() string {
	return field.Name
}
