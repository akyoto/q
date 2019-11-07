package build

import "github.com/akyoto/q/spec"

// Variable represents both local variables and function parameters.
type Variable struct {
	Name      string
	Type      *spec.Type
	Mutable   bool
	TimesUsed int
	Register  string
}

// String returns the string representation.
func (variable *Variable) String() string {
	return variable.Name
}
