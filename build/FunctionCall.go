package build

// FunctionCall represents a function call in the source code.
type FunctionCall struct {
	Function   *Function
	Parameters []Expression
}
