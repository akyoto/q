package ssa

import (
	"fmt"

	"git.urbach.dev/cli/q/src/types"
)

// Syscall temporarily transfers program flow to the OS kernel.
type Syscall struct {
	Arguments
	Liveness
	Source
}

// Equals returns true if the syscalls are equal.
func (a *Syscall) Equals(v Value) bool {
	b, sameType := v.(*Syscall)

	if !sameType {
		return false
	}

	return a.Arguments.Equals(b.Arguments)
}

// IsPure returns false because a syscall can have side effects.
func (s *Syscall) IsPure() bool { return false }

// String returns a human-readable representation of the syscall.
func (s *Syscall) String() string { return fmt.Sprintf("syscall(%s)", s.Arguments.String()) }

// Type returns the type of the syscall.
func (s *Syscall) Type() types.Type { return types.Any }