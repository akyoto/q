package instruction

// Instruction encapsulates a single instruction inside a function.
// Instructions can be variable assignments, function calls or keywords.
type Instruction struct {
	Expression Expression
	Kind       Kind
}

// Kind represents the type of an instruction.
type Kind uint8

const (
	// Unknown represents an invalid instruction.
	Unknown Kind = iota

	// Assignment moves data inside a variable or struct field.
	Assignment

	// Call represents a function call.
	Call

	// Keyword represents an instruction based on a keyword.
	Keyword
)
