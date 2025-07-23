package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

type Field struct {
	Object Value
	Field  *types.Field
	Liveness
	Source
}

func (v *Field) IsConst() bool        { return true }
func (v *Field) Type() types.Type     { return v.Field.Type }
func (v *Field) Replace(Value, Value) {}
func (v *Field) String() string       { return fmt.Sprintf("%s.%s", v.Object, v.Field) }
func (v *Field) Inputs() []Value      { return []Value{v.Object} }

func (a *Field) Equals(v Value) bool {
	b, sameType := v.(*Field)

	if !sameType {
		return false
	}

	return a.Field == b.Field
}