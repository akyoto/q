package types

import "github.com/akyoto/q/build/token"

// Field is a field in a data structure.
type Field struct {
	Name     string
	Type     *Type
	Mutable  bool
	Position token.Position
}

// String returns the name of the parameter.
func (field *Field) String() string {
	return field.Name
}
