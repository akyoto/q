package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

// Call is an internal function call.
type Call struct {
	Func *Function
	Arguments
	Liveness
	Source
}

// Equals returns true if the calls are equal.
func (a *Call) Equals(v Value) bool {
	b, sameType := v.(*Call)

	if !sameType {
		return false
	}

	return a.Arguments.Equals(b.Arguments)
}

// IsConst returns false because a function call can have side effects.
func (c *Call) IsConst() bool { return false }

// String returns a human-readable representation of the call.
func (c *Call) String() string {
	return fmt.Sprintf("%s(%s)", c.Func.String(), c.Arguments.String())
}

// Type returns the return type of the function.
func (c *Call) Type() types.Type {
	if len(c.Func.Typ.Output) == 0 {
		return types.Void
	}

	if len(c.Func.Typ.Output) == 1 {
		return c.Func.Typ.Output[0]
	}

	return &types.Tuple{Types: c.Func.Typ.Output}
}