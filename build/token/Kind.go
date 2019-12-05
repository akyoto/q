package token

// Kind represents the type of token.
type Kind uint8

const (
	// Invalid represents an invalid token.
	Invalid Kind = iota

	// NewLine represents the newline character.
	NewLine

	// Identifier represents a series of characters used to identify a variable or function.
	Identifier

	// Keyword represents a language keyword.
	Keyword

	// Text represents an uninterpreted series of characters in the source code.
	Text

	// Number represents a series of numerical characters.
	Number

	// Operator represents a mathematical operator.
	Operator

	// Separator represents a comma.
	Separator

	// Range represents '..'.
	Range

	// Question represents '?'.
	Question

	// Comment represents a comment.
	Comment

	// GroupStart represents '('.
	GroupStart

	// GroupEnd represents ')'.
	GroupEnd

	// BlockStart represents '{'.
	BlockStart

	// BlockEnd represents '}'.
	BlockEnd

	// ArrayStart represents '['.
	ArrayStart

	// ArrayEnd represents ']'.
	ArrayEnd
)

// String returns the text representation.
func (kind Kind) String() string {
	switch kind {
	case NewLine:
		return "NewLine"

	case Identifier:
		return "Identifier"

	case Keyword:
		return "Keyword"

	case Text:
		return "Text"

	case Number:
		return "Number"

	case Operator:
		return "Operator"

	case Separator:
		return "Separator"

	case Range:
		return "Range"

	case Question:
		return "Question"

	case Comment:
		return "Comment"

	case GroupStart:
		return "GroupStart"

	case GroupEnd:
		return "GroupEnd"

	case BlockStart:
		return "BlockStart"

	case BlockEnd:
		return "BlockEnd"

	case ArrayStart:
		return "ArrayStart"

	case ArrayEnd:
		return "ArrayEnd"

	case Invalid:
		return "Invalid"

	default:
		return "<undefined token>"
	}
}
