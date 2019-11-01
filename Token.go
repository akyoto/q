package main

type Token int

const (
	TokenIdentifier Token = iota
	TokenText
	TokenBracketStart
	TokenBracketEnd
)
