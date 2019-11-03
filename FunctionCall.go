package main

import "github.com/akyoto/q/spec"

// FunctionCall
type FunctionCall struct {
	Function        *spec.Function
	Parameters      []Expression
	ProcessedTokens int
}
