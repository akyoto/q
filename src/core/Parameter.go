package core

import (
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// Parameter is an input or output parameter in a function.
type Parameter struct {
	Name       string
	TypeTokens token.List
	typ        types.Type
}

// Type returns the data type of the parameter.
func (p *Parameter) Type() types.Type {
	return p.typ
}