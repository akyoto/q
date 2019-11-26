package expression

// Kind represents the type of an expression.
type Kind uint8

const (
	// Empty represents an empty expression.
	Empty Kind = iota

	// Operator represents an operation with operands.
	Operator

	// Call represents a function call.
	Call

	// Identifier represents a variable.
	Identifier

	// Number represents a number.
	Number

	// Text represents a string of characters.
	Text
)
