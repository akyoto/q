package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

// CallPointer is a call using a function pointer.
type CallPointer struct {
	Arguments
	Liveness
	Source
}

// Equals returns true if the calls are equal.
func (a *CallPointer) Equals(v Value) bool {
	b, sameType := v.(*CallPointer)

	if !sameType {
		return false
	}

	return a.Arguments.Equals(b.Arguments)
}

// IsPure returns false because a function call can have side effects.
func (c *CallPointer) IsPure() bool { return false }

// String returns a human-readable representation of the call.
func (c *CallPointer) String() string {
	return fmt.Sprintf("%p(%s)", c.Arguments[0], c.Arguments[1:].String())
}

// Type returns the return type of the function.
func (c *CallPointer) Type() types.Type {
	output := c.Arguments[0].Type().(*types.Function).Output

	if len(output) == 0 {
		return types.Void
	}

	if len(output) == 1 {
		return output[0]
	}

	tuple := &types.Tuple{}

	for _, typ := range output {
		structure, isStruct := types.Unwrap(typ).(*types.Struct)

		if !isStruct {
			tuple.Types = append(tuple.Types, typ)
			continue
		}

		for _, field := range structure.Fields {
			tuple.Types = append(tuple.Types, field.Type)
		}
	}

	return tuple
}