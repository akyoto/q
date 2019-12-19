package operators

// Operator represents an operator for mathematical expressions.
type Operator struct {
	Symbol                string
	Priority              uint8
	Kind                  Kind
	OperandOrderImportant bool
}
