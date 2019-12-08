package build

import "github.com/akyoto/q/build/token"

// Parameter represents a function parameter.
type Parameter struct {
	Name     string
	Type     *Type
	Mutable  bool
	Position token.Position
}

// String returns the name of the parameter.
func (parameter *Parameter) String() string {
	return parameter.Name
}
