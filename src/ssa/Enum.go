package ssa

import "git.urbach.dev/cli/q/src/types"

// Enum is an enum type reference.
type Enum struct {
	Typ *types.Enum
	Independent
	Void
	Source
}

// Equals returns true if the enums are equal.
func (a *Enum) Equals(v Value) bool {
	b, sameType := v.(*Enum)

	if !sameType {
		return false
	}

	return a.Typ == b.Typ
}

// String returns the name of the enum type.
func (e *Enum) String() string {
	return e.Typ.Name()
}