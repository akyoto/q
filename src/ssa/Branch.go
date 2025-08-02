package ssa

import (
	"fmt"
)

type Branch struct {
	Void
	Condition Value
	Then      *Block
	Else      *Block
}

func (a *Branch) Equals(v Value) bool {
	b, sameType := v.(*Branch)

	if !sameType {
		return false
	}

	return a.Condition.Equals(b.Condition) && a.Then == b.Then && a.Else == b.Else
}

func (v *Branch) Inputs() []Value {
	return []Value{v.Condition}
}

func (v *Branch) Replace(old Value, new Value) {
	if v.Condition == old {
		v.Condition = new
	}
}

func (v *Branch) String() string {
	return fmt.Sprintf("branch(%s, %s, %s)", v.Condition.String(), v.Then.Label, v.Else.Label)
}