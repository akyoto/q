package ssa

// Block is a list of instructions that can be targeted in branches.
type Block struct {
	Instructions []Value
}

// Append adds a new instruction to the block.
func (b *Block) Append(instr Value) *Value {
	b.Instructions = append(b.Instructions, instr)
	return &b.Instructions[len(b.Instructions)-1]
}