package build

import "github.com/akyoto/q/build/spec"

// Parameter represents a function parameter.
type Parameter struct {
	Name    string
	Type    *spec.Type
	Mutable bool
}

// String returns the name of the parameter.
func (parameter *Parameter) String() string {
	return parameter.Name
}
