package build

import (
	"fmt"

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
	NoParameterCheck bool
	SideEffects      int32
	CallCount        int32
	parameterStart   token.Position
}

// Tokens returns all tokens within the function body (excluding the braces '{' and '}').
func (function *Function) Tokens() []token.Token {
	return function.File.tokens[function.TokenStart:function.TokenEnd]
}

// Error creates an error inside the function.
func (function *Function) Error(position token.Position, err error) error {
	metaError, hasMetaData := err.(*Error)

	if hasMetaData {
		return metaError
	}

	return NewError(err, function.File.path, function.File.tokens[:function.TokenStart+position+1])
}

// Errorf creates a formatted error inside the function.
func (function *Function) Errorf(position token.Position, message string, args ...interface{}) error {
	return function.Error(position, fmt.Errorf(message, args...))
}
