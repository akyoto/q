package token

// Kind represents the type of token.
type Kind uint8

const (
	// Unknown represents an invalid token.
	Unknown Kind = iota

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

	// GroupStart represents '('.
	GroupStart

	// GroupEnd represents ')'.
	GroupEnd

	// BlockStart represents '{'.
	BlockStart

	// BlockEnd represents '}'.
	BlockEnd
)

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

	case GroupStart:
		return "GroupStart"

	case GroupEnd:
		return "GroupEnd"

	case BlockStart:
		return "BlockStart"

	case BlockEnd:
		return "BlockEnd"

	case Unknown:
		return "Unknown"

	default:
		return "<undefined>"
	}
}
