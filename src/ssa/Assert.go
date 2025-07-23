package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

type Assert struct {
	Condition Value
	Source
}

func (v *Assert) AddUser(Value)    { panic("assert can not be used as a dependency") }
func (v *Assert) Inputs() []Value  { return []Value{v.Condition} }
func (v *Assert) IsConst() bool    { return false }
func (v *Assert) Type() types.Type { return types.Void }
func (v *Assert) Users() []Value   { return nil }

func (a *Assert) Equals(v Value) bool {
	b, sameType := v.(*Assert)

	if !sameType {
		return false
	}

	return a.Condition.Equals(b.Condition)
}

func (v *Assert) Replace(old Value, new Value) {
	if v.Condition == old {
		v.Condition = new
	}
}

func (v *Assert) String() string {
	return fmt.Sprintf("assert %s", v.Condition.String())
}