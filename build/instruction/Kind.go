package instruction

// Kind represents the type of an instruction.
type Kind uint8

const (
	// Invalid represents an invalid instruction.
	Invalid Kind = iota

	// Assignment moves data inside a variable or struct field.
	Assignment

	// Call represents a function call.
	Call

	// Keyword represents an instruction based on a keyword.
	Keyword
)

// String returns the text representation.
func (kind Kind) String() string {
	switch kind {
	case Invalid:
		return "Invalid"

	case Assignment:
		return "Assignment"

	case Call:
		return "Call"

	case Keyword:
		return "Keyword"

	default:
		return "<undefined instruction>"
	}
}
