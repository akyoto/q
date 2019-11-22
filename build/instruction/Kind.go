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

	// IfStart represents the start of the branch.
	IfStart

	// IfEnd represents the end of the branch.
	IfEnd

	// LoopStart represents the start of the loop.
	LoopStart

	// LoopEnd represents the end of the loop.
	LoopEnd

	// Return represents the return statement.
	Return
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

	case LoopStart:
		return "LoopStart"

	case LoopEnd:
		return "LoopEnd"

	default:
		return "<undefined instruction>"
	}
}
