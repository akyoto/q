package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

type StructField struct {
	Struct Value
	Field  *types.Field
	Id
	Liveness
	Source
}

func (v *StructField) Dependencies() []Value {
	return []Value{v.Struct}
}

func (a *StructField) Equals(v Value) bool {
	b, sameType := v.(*StructField)

	if !sameType {
		return false
	}

	return a.Field == b.Field
}

func (v *StructField) IsConst() bool {
	return true
}

func (v *StructField) Debug(expand bool) string {
	if expand {
		return fmt.Sprintf("%s.%s", v.Struct, v.Field)
	}

	return fmt.Sprintf("%%%d.%s", v.Struct.ID(), v.Field)
}

func (v *StructField) String() string {
	return v.Debug(true)
}

func (v *StructField) Type() types.Type {
	return v.Field.Type
}