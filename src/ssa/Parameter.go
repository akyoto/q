package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

// Parameter is an input parameter for a function call.
type Parameter struct {
	Typ       types.Type
	Name      string
	Tokens    token.List
	Structure *Struct
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

// Inputs returns nil because a parameter has no inputs.
func (p *Parameter) Inputs() []Value { return nil }

// IsConst returns true because a parameter is always constant.
func (p *Parameter) IsConst() bool { return true }

// Replace does nothing because a parameter has no inputs.
func (p *Parameter) Replace(Value, Value) {}

// Struct returns the struct that this parameter is a part of.
func (p *Parameter) Struct() *Struct { return p.Structure }

// String returns a human-readable representation of the parameter.
func (p *Parameter) String() string { return fmt.Sprintf("args[%d]", p.Index) }

// Type returns the type of the parameter.
func (p *Parameter) Type() types.Type { return p.Typ }