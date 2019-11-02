package token

type Kind int

const (
	Unknown Kind = iota
	StartOfLine
	Identifier
	Keyword
	Text
	ParenthesesStart
	ParenthesesEnd
	BlockStart
	BlockEnd
	WhiteSpace
)
