package main

type TokenKind int

const (
	TokenStartOfLine TokenKind = iota
	TokenIdentifier
	TokenText
	TokenBracketStart
	TokenBracketEnd
	TokenWhiteSpace
)

type Token struct {
	Kind TokenKind
	Text []byte
}
