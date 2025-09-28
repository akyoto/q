package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/cpu"
	"git.urbach.dev/cli/q/src/types"
)

// Register represents a CPU register.
type Register struct {
	Register cpu.Register
	Liveness
	Source
}

// Equals returns true if the syscalls are equal.
func (a *Register) Equals(v Value) bool {
	b, sameType := v.(*Register)

	if !sameType {
		return false
	}

	return a.Register == b.Register
}

// Inputs returns nil.
func (s *Register) Inputs() []Value { return nil }

// IsPure returns true because a register access has no side effects.
func (s *Register) IsPure() bool { return true }

// Replace does nothing.
func (s *Register) Replace(Value, Value) {}

// String returns a human-readable representation of the register.
func (s *Register) String() string { return fmt.Sprintf("register(%s)", s.Register.String()) }

// Type returns the type of the register.
func (s *Register) Type() types.Type { return types.Any }