package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

// Phi is the merging point of multiple values for the same name.
type Phi struct {
	Typ types.Type
	Arguments
	Liveness
}

// DefinedArguments is an iterator over arguments that are defined.
func (p *Phi) DefinedArguments(yield func(Value) bool) {
	for _, arg := range p.Arguments {
		if arg == Undefined {
			continue
		}

		if !yield(arg) {
			return
		}
	}
}

// Equals returns true if the phi nodes are equal.
func (a *Phi) Equals(v Value) bool {
	b, sameType := v.(*Phi)

	if !sameType {
		return false
	}

	return a.Arguments.Equals(b.Arguments)
}

// FirstDefined returns the first input value that is not undefined.
func (p *Phi) FirstDefined() Value {
	for _, arg := range p.Arguments {
		if arg != Undefined {
			return arg
		}
	}

	panic("phi composed of only undefined values must not exist")
}

// IsConst returns true because two equal phi nodes produce the same value.
func (p *Phi) IsConst() bool {
	return true
}

// IsPartiallyUndefined returns true if one of the input values is undefined.
func (p *Phi) IsPartiallyUndefined() bool {
	for _, arg := range p.Arguments {
		if arg == Undefined {
			return true
		}
	}

	return false
}

// String returns a human-readable representation of the phi node.
func (p *Phi) String() string {
	return fmt.Sprintf("phi(%s)", p.Arguments.String())
}

// Type returns the type of the phi node.
func (p *Phi) Type() types.Type {
	return p.Typ
}