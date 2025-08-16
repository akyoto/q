package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

// FromTuple is a value inside of a tuple.
type FromTuple struct {
	Tuple Value
	Index int
	Liveness
	Source
}

func (v *FromTuple) Inputs() []Value  { return []Value{v.Tuple} }
func (v *FromTuple) IsConst() bool    { return true }
func (v *FromTuple) String() string   { return fmt.Sprintf("%p [%d]", v.Tuple, v.Index) }
func (v *FromTuple) Type() types.Type { return v.Tuple.Type().(*types.Tuple).Types[v.Index] }

func (a *FromTuple) Equals(v Value) bool {
	b, sameType := v.(*FromTuple)

	if !sameType {
		return false
	}

	return a.Tuple == b.Tuple && a.Index == b.Index
}

func (v *FromTuple) Replace(old Value, new Value) {
	if v.Tuple == old {
		v.Tuple = new
	}
}