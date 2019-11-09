package build

import (
	"github.com/akyoto/q/instruction"
	"github.com/akyoto/q/spec"
	"github.com/akyoto/q/token"
)

// Function represents a function.
type Function struct {
	Name             string
	Parameters       []Variable
	ReturnTypes      []spec.Type
	TokenStart       token.Position
	TokenEnd         token.Position
	File             *File
	NoParameterCheck bool
	parameterStart   token.Position
}

// Tokens returns all tokens within the function body (excluding the braces '{' and '}').
func (function *Function) Tokens() []token.Token {
	return function.File.tokens[function.TokenStart:function.TokenEnd]
}

// Instructions returns all instructions within the function body.
func (function *Function) Instructions() []instruction.Instruction {
	return instruction.FromTokens(function.Tokens())
}

// Error creates an error inside the function.
func (function *Function) Error(message string, position token.Position) error {
	return NewError(message, function.File.path, function.File.tokens[:function.TokenStart+position+1])
}
