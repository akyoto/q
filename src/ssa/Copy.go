package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

// Copy is an operation that clones a value.
type Copy struct {
	Value Value
	Liveness
	Source
}

func (a *Copy) Equals(v Value) bool {
	b, sameType := v.(*Copy)

	if !sameType {
		return false
	}

	return a.Value == b.Value
}

func (v *Copy) Inputs() []Value {
	return []Value{v.Value}
}

func (v *Copy) IsConst() bool {
	return false
}

func (v *Copy) Replace(old Value, new Value) {
	if v.Value == old {
		v.Value = new
	}
}

func (v *Copy) String() string {
	return fmt.Sprintf("copy(%p)", v.Value)
}

func (v *Copy) Type() types.Type {
	return v.Value.Type()
}