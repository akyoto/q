package spec

// Variable represents both local variables and function parameters.
type Variable struct {
	Name string
	Type *Type
}

// String returns the string representation.
func (variable *Variable) String() string {
	return variable.Name
}
