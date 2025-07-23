package ssa

import (
	"strings"

	"git.urbach.dev/cli/q/src/token"
)

// Arguments defines a list of values that this value depends on.
type Arguments []Value

func (v Arguments) Inputs() []Value { return v }

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

func (v Arguments) Replace(old Value, new Value) {
	for i, arg := range v {
		if arg == old {
			v[i] = new
		}
	}
}

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

// Liveness tracks where the value is used.
type Liveness struct {
	users []Value
}

func (v *Liveness) AddUser(user Value) { v.users = append(v.users, user) }
func (v *Liveness) ReplaceUser(old Value, new Value) {
	for i, user := range v.users {
		if user == old {
			v.users[i] = new
		}
	}
}
func (v *Liveness) Users() []Value { return v.users }

type HasLiveness interface {
	AddUser(Value)
	ReplaceUser(Value, Value)
	Users() []Value
}

// Source tracks the source tokens.
type Source token.Source

func (v Source) Start() token.Position         { return v.StartPos }
func (v Source) End() token.Position           { return v.EndPos }
func (v Source) StringFrom(code []byte) string { return string(code[v.Start():v.End()]) }

type HasSource interface {
	Start() token.Position
	End() token.Position
	StringFrom([]byte) string
}

type StructField interface {
	Struct() *Struct
}