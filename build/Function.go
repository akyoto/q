package build

import (
	"fmt"

	"github.com/akyoto/q/build/instruction"
	"github.com/akyoto/q/build/spec"
	"github.com/akyoto/q/build/token"
)

// Function represents a function.
type Function struct {
	Name             string
	Parameters       []Variable
	ReturnTypes      []spec.Type
	File             *File
	TokenStart       token.Position
	TokenEnd         token.Position
	Used             bool
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
func (function *Function) Error(position token.Position, err error) error {
	return NewError(err, function.File.path, function.File.tokens[:function.TokenStart+position+1])
}

// Errorf creates a formatted error inside the function.
func (function *Function) Errorf(position token.Position, message string, args ...interface{}) error {
	return function.Error(position, fmt.Errorf(message, args...))
}
