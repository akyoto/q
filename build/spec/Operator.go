package spec

// Operator represents an operator for mathematical expressions.
type Operator struct {
	Symbol                string
	Priority              int
	OperandOrderImportant bool
}
