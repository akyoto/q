package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

// Struct is a list of values that can be addressed by name.
type Struct struct {
	Typ types.Type
	Arguments
	Liveness
	Source
}

// Equals returns true if the structs are equal.
func (a *Struct) Equals(v Value) bool {
	b, sameType := v.(*Struct)

	if !sameType {
		return false
	}

	return a.Arguments.Equals(b.Arguments)
}

// IsConst returns true because two instances using the same values are interchangeable.
func (s *Struct) IsConst() bool { return true }

// String returns a human-readable representation of the struct.
func (s *Struct) String() string {
	return fmt.Sprintf("%s{%s}", s.Typ.Name(), s.Arguments.String())
}

// Type returns the type of the struct.
func (s *Struct) Type() types.Type { return s.Typ }