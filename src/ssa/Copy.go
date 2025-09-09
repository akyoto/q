package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

// Copy is an operation that clones a value.
type Copy struct {
	Value Value
	Typ   types.Type
	Liveness
	Source
}

// Equals returns true if the copies are equal.
func (a *Copy) Equals(v Value) bool {
	b, sameType := v.(*Copy)

	if !sameType {
		return false
	}

	return a.Value == b.Value
}

// Inputs returns the value to be copied.
func (c *Copy) Inputs() []Value {
	return []Value{c.Value}
}

// IsConst returns true.
func (c *Copy) IsConst() bool {
	return true
}

// Replace replaces the value to be copied if it matches.
func (c *Copy) Replace(old Value, new Value) {
	if c.Value == old {
		c.Value = new
	}
}

// String returns a human-readable representation of the copy.
func (c *Copy) String() string {
	if c.Typ != c.Value.Type() {
		return fmt.Sprintf("copy(%p as %s)", c.Value, c.Typ.Name())
	}

	return fmt.Sprintf("copy(%p)", c.Value)
}

// Type returns the type of the copied value.
func (c *Copy) Type() types.Type {
	return c.Typ
}