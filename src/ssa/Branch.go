package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

type Branch struct {
	Condition Value
	Then      *Block
	Else      *Block
	Source
}

func (v *Branch) AddUser(Value)    { panic("if can not be used as a dependency") }
func (v *Branch) Inputs() []Value  { return []Value{v.Condition} }
func (v *Branch) IsConst() bool    { return false }
func (v *Branch) Type() types.Type { return types.Void }
func (v *Branch) Users() []Value   { return nil }

func (a *Branch) Equals(v Value) bool {
	b, sameType := v.(*Branch)

	if !sameType {
		return false
	}

	return a.Condition.Equals(b.Condition) && a.Then == b.Then && a.Else == b.Else
}

func (v *Branch) String() string {
	return fmt.Sprintf("branch(%s, %s, %s)", v.Condition.String(), v.Then.Label, v.Else.Label)
}