package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

type Field struct {
	Struct Value
	Field  *types.Field
	Liveness
	Source
}

func (v *Field) IsConst() bool    { return true }
func (v *Field) Type() types.Type { return v.Field.Type }
func (v *Field) String() string   { return fmt.Sprintf("%s.%s", v.Struct, v.Field) }
func (v *Field) Inputs() []Value  { return []Value{v.Struct} }

func (a *Field) Equals(v Value) bool {
	b, sameType := v.(*Field)

	if !sameType {
		return false
	}

	return a.Field == b.Field
}