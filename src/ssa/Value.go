package ssa

import (
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

type Value interface {
	AddUse(Value)
	Alive() int
	Dependencies() []Value
	End() token.Position
	Equals(Value) bool
	IsConst() bool
	String() string
	Start() token.Position
	Type() types.Type
}