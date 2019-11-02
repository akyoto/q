package main

type TokenKind int

const (
	TokenIdentifier TokenKind = iota
	TokenText
	TokenBracketStart
	TokenBracketEnd
	TokenEndOfLine
)

type Token struct {
	Kind TokenKind
	Text []byte
}
