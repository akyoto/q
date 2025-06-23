package ssa

import "git.urbach.dev/cli/q/src/token"

type Value interface {
	AddUse(Value)
	Alive() int
	Dependencies() []Value
	Equals(Value) bool
	IsConst() bool
	String() string
	Token() token.Token
}