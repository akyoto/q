package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

type Parameter struct {
	Index  uint8
	Name   string
	Typ    types.Type
	Tokens token.List
	Liveness
}

func (v *Parameter) Inputs() []Value  { return nil }
func (v *Parameter) IsConst() bool    { return true }
func (v *Parameter) String() string   { return fmt.Sprintf("args[%d]", v.Index) }
func (v *Parameter) Type() types.Type { return v.Typ }

func (a *Parameter) Equals(v Value) bool {
	b, sameType := v.(*Parameter)

	if !sameType {
		return false
	}

	return a.Index == b.Index
}