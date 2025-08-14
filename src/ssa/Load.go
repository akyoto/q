package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

// Load stores a value at a given index relative to the address.
type Load struct {
	Address Value
	Index   Value
	Liveness
	Source
}

func (a *Load) Equals(v Value) bool {
	b, sameType := v.(*Load)

	if !sameType {
		return false
	}

	return a.Address == b.Address && a.Index == b.Index
}

func (v *Load) IsConst() bool {
	return false
}

func (v *Load) Inputs() []Value {
	return []Value{v.Address, v.Index}
}

func (v *Load) Replace(old Value, new Value) {
	if v.Address == old {
		v.Address = new
	}

	if v.Index == old {
		v.Index = new
	}
}

func (v *Load) String() string {
	return fmt.Sprintf("load(%p, %p)", v.Address, v.Index)
}

func (v *Load) Type() types.Type {
	return v.Address.Type().(*types.Pointer).To
}