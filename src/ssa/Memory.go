package ssa

import "git.urbach.dev/cli/q/src/types"

// Memory represents a memory address.
type Memory struct {
	Typ     types.Type
	Address Value
	Index   Value
	Scale   bool
}