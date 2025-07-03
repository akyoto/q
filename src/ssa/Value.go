package ssa

import (
	"git.urbach.dev/cli/q/src/token"
	"git.urbach.dev/cli/q/src/types"
)

type Value interface {
	// Essentials
	Debug(bool) string
	ID() int
	IsConst() bool
	SetID(int)
	String() string
	Type() types.Type

	// Arguments
	Dependencies() []Value
	Equals(Value) bool

	// Liveness
	AddUser(Value)
	CountUsers() int

	// Source
	SetSource(token.List)
	Start() token.Position
	End() token.Position
}