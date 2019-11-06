package token

// Kind represents the type of token.
type Kind uint8

const (
	Unknown Kind = iota
	NewLine
	Identifier
	Keyword
	Text
	Number
	GroupStart
	GroupEnd
	BlockStart
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

	case GroupStart:
		return "GroupStart"

	case GroupEnd:
		return "GroupEnd"

	case BlockStart:
		return "BlockStart"

	case BlockEnd:
		return "BlockEnd"

	default:
		return "Unknown"
	}
}
