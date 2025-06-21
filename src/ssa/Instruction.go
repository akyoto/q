package ssa

import (
	"fmt"
)

// Instruction is a single instruction in a basic block.
// It is implemented as a "fat struct" for performance reasons.
// It contains all the fields necessary to represent all instruction types.
type Instruction struct {
	Args []*Instruction
	Int  int64
	Type Type
}

// String returns a human-readable representation of the instruction.
func (i *Instruction) String() string {
	switch i.Type {
	case Int:
		return fmt.Sprintf("%d", i.Int)
	case Add:
		return fmt.Sprintf("%s + %s", i.Args[0], i.Args[1])
	default:
		return ""
	}
}