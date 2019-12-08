package build

// Parameter represents a function parameter.
type Parameter struct {
	Name    string
	Type    *Type
	Mutable bool
}

// String returns the name of the parameter.
func (parameter *Parameter) String() string {
	return parameter.Name
}
