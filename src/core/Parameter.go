package core

import (
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// Parameter is an input or output parameter in a function.
type Parameter struct {
	name   string
	typ    types.Type
	tokens token.List
}

// NewParameter creates a new parameter with the given list of tokens.
func NewParameter(tokens token.List) *Parameter {
	return &Parameter{tokens: tokens}
}

// Name returns the name of the parameter.
func (p *Parameter) Name() string {
	return p.name
}

// Type returns the type of the parameter.
func (p *Parameter) Type() types.Type {
	return p.typ
}