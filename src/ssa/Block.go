package ssa

import (
	"slices"
)

// Block is a list of instructions that can be targeted in branches.
type Block struct {
	Label        string
	Instructions []Value
}

// Append adds a new instruction to the block.
func (block *Block) Append(instr Value) Value {
	for _, dep := range instr.Inputs() {
		dep.(HasLiveness).AddUser(instr)
	}

	block.Instructions = append(block.Instructions, instr)
	return instr
}

// RemoveNilValues removes all nil values from the block.
func (block *Block) RemoveNilValues() {
	block.Instructions = slices.DeleteFunc(block.Instructions, func(value Value) bool {
		return value == nil
	})
}