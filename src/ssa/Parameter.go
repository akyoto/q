package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// Parameter is an input parameter for a function call.
type Parameter struct {
	Typ    types.Type
	Name   string
	Tokens token.List
	Independent
	Liveness
	Source
	Index uint8
}

// Equals returns true if the parameters are equal.
func (a *Parameter) Equals(v Value) bool {
	b, sameType := v.(*Parameter)

	if !sameType {
		return false
	}

	return a.Index == b.Index
}

// IsPure returns true because a parameter is always constant.
func (p *Parameter) IsPure() bool { return true }

// String returns a human-readable representation of the parameter.
func (p *Parameter) String() string { return fmt.Sprintf("args[%d]", p.Index) }

// Type returns the type of the parameter.
func (p *Parameter) Type() types.Type { return p.Typ }