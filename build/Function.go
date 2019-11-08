package build

import (
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
