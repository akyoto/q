package ssa

import (
	"strings"

	"git.urbach.dev/cli/q/src/token"
)

// Arguments defines a list of values that this value depends on.
type Arguments []Value

func (v Arguments) Inputs() []Value { return v }

// String returns a comma-separated list of all arguments.
func (v Arguments) String() string {
	tmp := strings.Builder{}

	for i, arg := range v {
		tmp.WriteString(arg.String())

		if i != len(v)-1 {
			tmp.WriteString(", ")
		}
	}

	return tmp.String()
}

func (a Arguments) Equals(b Arguments) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if !a[i].Equals(b[i]) {
			return false
		}
	}

	return true
}

// Liveness tracks where the value is used.
type Liveness struct {
	users []Value
}

func (v *Liveness) AddUser(user Value) { v.users = append(v.users, user) }
func (v *Liveness) Users() []Value     { return v.users }

type HasLiveness interface {
	AddUser(Value)
	Users() []Value
}

// Source tracks the source tokens.
type Source token.List

func (v Source) Start() token.Position { return v[0].Position }
func (v Source) End() token.Position   { return v[len(v)-1].End() }

type HasSource interface {
	Start() token.Position
	End() token.Position
}