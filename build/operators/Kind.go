package operators

// Kind represents the type of an operator.
type Kind uint8

const (
	// Default represents the default type.
	Default Kind = iota

	// Assignment moves data inside a variable or struct field.
	Assignment

	// Call represents a function call.
	Comparison
)

// String returns the text representation.
func (kind Kind) String() string {
	switch kind {
	case Assignment:
		return "Assignment"

	case Comparison:
		return "Comparison"

	case Default:
		return "Default"

	default:
		return "<undefined operator>"
	}
}
