package token

// Kind represents the type of token.
type Kind int

const (
	Unknown Kind = iota
	StartOfLine
	Identifier
	Keyword
	Text
	GroupStart
	GroupEnd
	BlockStart
	BlockEnd
	WhiteSpace
)
