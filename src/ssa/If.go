package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

type If struct {
	Condition Value
	Then      *Block
	Else      *Block
	Source
}

func (v *If) AddUser(Value)    { panic("if can not be used as a dependency") }
func (v *If) Inputs() []Value  { return []Value{v.Condition} }
func (v *If) IsConst() bool    { return false }
func (v *If) Type() types.Type { return types.Void }
func (v *If) Users() []Value   { return nil }

func (a *If) Equals(v Value) bool {
	b, sameType := v.(*If)

	if !sameType {
		return false
	}

	return a.Condition.Equals(b.Condition) && a.Then == b.Then && a.Else == b.Else
}

func (v *If) String() string {
	return fmt.Sprintf("if %s then %v else %v", v.Condition.String(), v.Then, v.Else)
}