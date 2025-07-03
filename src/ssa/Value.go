package ssa

import (
	"git.urbach.dev/cli/q/src/types"
)

type Value interface {
	Debug(bool) string
	Dependencies() []Value
	Equals(Value) bool
	ID() int
	IsConst() bool
	SetID(int)
	String() string
	Type() types.Type
}