package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

type Jump struct {
	To *Block
	Source
}

func (v *Jump) AddUser(Value)        { panic("jump can not be used as a dependency") }
func (v *Jump) Inputs() []Value      { return nil }
func (v *Jump) IsConst() bool        { return false }
func (v *Jump) Replace(Value, Value) {}
func (v *Jump) Type() types.Type     { return types.Void }
func (v *Jump) Users() []Value       { return nil }

func (a *Jump) Equals(v Value) bool {
	b, sameType := v.(*Jump)

	if !sameType {
		return false
	}

	return a.To == b.To
}

func (v *Jump) String() string {
	return fmt.Sprintf("jump(%s)", v.To.Label)
}